package services

import (
	"github.com/danielgerlag/goflow/interfaces"
)

// WorkflowExecutor runs workflows
type WorkflowExecutor struct {
	host interfaces.WorkflowHost
}

func NewWorkflowExecutor(host interfaces.WorkflowHost) *WorkflowExecutor {
	return &WorkflowExecutor{
		host: host,
	}
}

//Start starts the host
func (executor WorkflowExecutor) Start() {

}
