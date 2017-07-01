package models

type WorkflowDefinition struct {
	ID      string
	Version int
	Steps   []WorkflowStep
}
