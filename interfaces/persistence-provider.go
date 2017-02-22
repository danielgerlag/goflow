package interfaces

import (
	"github.com/danielgerlag/goflow/models"
)

type PersistenceProvider interface {
	CreateNewWorkflow(workflow models.WorkflowInstance) string
	PersistWorkflow(workflow models.WorkflowInstance)
	GetRunnableInstances() []string
	GetWorkflowInstance(id string) models.WorkflowInstance
}
