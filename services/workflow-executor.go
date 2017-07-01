package services

import (
	"reflect"

	linq "gopkg.in/ahmetb/go-linq.v3"

	"github.com/danielgerlag/goflow/interfaces"
	"github.com/danielgerlag/goflow/models"
)

// WorkflowExecutor runs workflows
type workflowExecutor struct {
	host     interfaces.WorkflowHost
	registry interfaces.WorkflowRegistry
}

func NewWorkflowExecutor(host interfaces.WorkflowHost) interfaces.WorkflowExecutor {

	return &workflowExecutor{
		host: host,
	}
}

func (executor workflowExecutor) Execute(instance models.WorkflowInstance) {

	var def = executor.registry.GetDefinition(instance.WorkflowDefinitionID, instance.Version)

	var activePointers []models.ExecutionPointer

	linq.From(instance.ExecutionPointers).WhereT(func(x models.ExecutionPointer) bool {
		return x.Active == true
	}).ToSlice(&activePointers)

	for _, pointer := range activePointers {
		step := linq.From(def.Steps).FirstWithT(func(x models.WorkflowStep) bool {
			return x.ID == pointer.StepID
		}).(models.WorkflowStep)

		//if pointer.StartTime == nil {
		//	pointer.StartTime = time.Now()
		//}

		var body = reflect.New(step.Body)

		//&body.
		//body.

	}
}
