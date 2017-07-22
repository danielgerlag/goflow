package services

import (
	"github.com/danielgerlag/goflow/models"
)

type workflowBuilder struct {
}

func (builder *workflowBuilder) Builder(id string, version int) models.WorkflowDefinition {

}

func (builder *workflowBuilder) LastStep() int {

}

func (builder *workflowBuilder) AddStep(step models.WorkflowStep) {

}

//instanceY := container.Instance(reflect.TypeOf((Y)(nil)))
