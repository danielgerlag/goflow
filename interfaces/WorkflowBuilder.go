package interfaces

import (
	"reflect"

	"github.com/danielgerlag/goflow/models"
)

type WorkflowBuilder interface {
	Build(id string, version int) models.WorkflowDefinition
	LastStep() int
	AddStep(step models.WorkflowStep)

	StartWith(step reflect.Type)
}
