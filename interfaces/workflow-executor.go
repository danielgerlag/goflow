package interfaces

import (
	"github.com/danielgerlag/goflow/models"
)

type WorkflowExecutor interface {
	Execute(instance *models.WorkflowInstance)
}
