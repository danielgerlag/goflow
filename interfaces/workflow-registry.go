package interfaces

import (
	"github.com/danielgerlag/goflow/models"
)

type WorkflowRegistry interface {
	GetDefinition(workflowID string, version int) (*models.WorkflowDefinition, error)
}
