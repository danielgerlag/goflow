package services

import (
	"reflect"
	"time"

	linq "gopkg.in/ahmetb/go-linq.v3"

	"github.com/danielgerlag/goflow/interfaces"
	"github.com/danielgerlag/goflow/models"
	"github.com/nu7hatch/gouuid"
)

// WorkflowExecutor runs workflows
type workflowExecutor struct {
	registry interfaces.WorkflowRegistry
}

func MakeWorkflowExecutor(registry interfaces.WorkflowRegistry) interfaces.WorkflowExecutor {

	return workflowExecutor{
		registry: registry,
	}
}

func (executor workflowExecutor) Execute(instance *models.WorkflowInstance) {

	var def = executor.registry.GetDefinition(instance.WorkflowDefinitionID, instance.Version)

	var activePointers []models.ExecutionPointer

	linq.From(instance.ExecutionPointers).WhereT(func(x models.ExecutionPointer) bool {
		return x.Active == true
	}).ToSlice(&activePointers)

	for _, pointer := range activePointers {
		step := linq.From(def.Steps).FirstWithT(func(x models.WorkflowStep) bool {
			return x.ID == pointer.StepID
		}).(models.WorkflowStep)

		if pointer.StartTime.IsZero() {
			pointer.StartTime = time.Now()
		}

		context := models.StepExectuionContext{}
		context.Step = step
		context.Workflow = *instance
		context.PersistenceData = pointer.PersistenceData

		body := reflect.New(step.Body).Elem().Interface().(interfaces.StepBody)

		stepResult := body.Run(context)
		processStepResult(stepResult, &pointer, instance, step)

	}

	determineNextExecution(instance)
}

func processStepResult(result models.ExecutionResult, pointer *models.ExecutionPointer, workflow *models.WorkflowInstance, step models.WorkflowStep) {
	pointer.PersistenceData = result.PersistenceData

	if result.Proceed {
		pointer.Active = false
		pointer.EndTime = time.Now()

		for _, outcome := range step.Outcomes {
			var u, _ = uuid.NewV4()

			newPointer := models.ExecutionPointer{}
			newPointer.ID = u.String()
			newPointer.Active = true
			newPointer.PredecessorID = pointer.PredecessorID
			newPointer.StepID = outcome.NextStep

			workflow.ExecutionPointers = append(workflow.ExecutionPointers, newPointer)
		}
	} else {
		for _, branch := range result.BranchValues {

			for _, child := range step.Children {
				var u, _ = uuid.NewV4()

				newPointer := models.ExecutionPointer{}
				newPointer.ID = u.String()
				newPointer.Active = true
				newPointer.PredecessorID = pointer.PredecessorID
				newPointer.StepID = child
				newPointer.ContextItem = branch

				workflow.ExecutionPointers = append(workflow.ExecutionPointers, newPointer)
				pointer.Children = append(pointer.Children, newPointer.ID)
			}
		}
	}

}

func determineNextExecution(workflow *models.WorkflowInstance) {
	workflow.NextExecution = time.Time{}

	if workflow.Status == models.Complete {
		return
	}

	var activePointersNoChildren []models.ExecutionPointer
	var activePointersWithChildren []models.ExecutionPointer

	linq.From(workflow.ExecutionPointers).WhereT(func(x models.ExecutionPointer) bool {
		return x.Active == true && len(x.Children) == 0
	}).ToSlice(&activePointersNoChildren)

	linq.From(workflow.ExecutionPointers).WhereT(func(x models.ExecutionPointer) bool {
		return x.Active == true && len(x.Children) > 0
	}).ToSlice(&activePointersWithChildren)

	for _, pointer := range activePointersNoChildren {
		if !pointer.SleepUntil.IsZero() {
			workflow.NextExecution = time.Now()
			return
		}

		if pointer.SleepUntil.Before(workflow.NextExecution) || workflow.NextExecution.IsZero() {
			workflow.NextExecution = pointer.SleepUntil
		}
	}

	if workflow.NextExecution.IsZero() {
		for _, pointer := range activePointersWithChildren {

			branchDone := linq.From(workflow.ExecutionPointers).WhereT(func(x models.ExecutionPointer) bool {
				return linq.From(x.Children).Contains(pointer.ID)
			}).AllT(func(x models.ExecutionPointer) bool {
				return isBranchComplete(workflow.ExecutionPointers, x.ID)
			})

			if branchDone {
				if !pointer.SleepUntil.IsZero() {
					workflow.NextExecution = time.Now()
					return
				}

				if pointer.SleepUntil.Before(workflow.NextExecution) || workflow.NextExecution.IsZero() {
					workflow.NextExecution = pointer.SleepUntil
				}
			}
		}
	}

	if workflow.NextExecution.IsZero() && linq.From(workflow.ExecutionPointers).AllT(func(x models.ExecutionPointer) bool { return !x.EndTime.IsZero() }) {
		workflow.Status = models.Complete
		workflow.CompleteTime = time.Now()
	}
}

func isBranchComplete(pointers []models.ExecutionPointer, rootID string) bool {
	root := linq.From(pointers).FirstWithT(func(x models.ExecutionPointer) bool {
		return x.ID == rootID
	}).(models.ExecutionPointer)

	if root.EndTime.IsZero() {
		return false
	}

	var list []models.ExecutionPointer

	linq.From(pointers).WhereT(func(x models.ExecutionPointer) bool {
		return x.PredecessorID == rootID
	}).ToSlice(&list)

	var result = true

	for _, item := range list {
		result = result && isBranchComplete(pointers, item.ID)
	}

	return result
}
