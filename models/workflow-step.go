package models

import (
	"reflect"
)

type WorkflowStep struct {
	ID   int
	Name string

	Body reflect.Type

	Outcomes []StepOutcome
	Inputs   []DataMapping
	Outputs  []DataMapping
}

//BodyType  type
