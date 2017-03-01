package models

type WorkflowStep struct {
	ID       int
	Name     string
	Outcomes []StepOutcome
	Inputs   []DataMapping
	Outputs  []DataMapping
}

//BodyType  type
