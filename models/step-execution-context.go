package models

type StepExectuionContext struct {
	Workflow        WorkflowInstance
	Step            WorkflowStep
	PersistenceData interface{}
}
