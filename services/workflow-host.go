package services

import (
	"github.com/danielgerlag/goflow/interfaces"
)

// WorkflowHost runs workflows
type WorkflowHost struct {
	workflowConsumer WorkflowConsumer
	runPoller        RunPoller
	persistence      interfaces.PersistenceProvider
}

func MakeHost(workflowConsumer WorkflowConsumer, runPoller RunPoller, persistence interfaces.PersistenceProvider) WorkflowHost {
	return WorkflowHost{
		workflowConsumer: workflowConsumer,
		runPoller:        runPoller,
		persistence:      persistence,
	}
}

//Start starts the host
func (host WorkflowHost) Start() {
	host.workflowConsumer.Start()
	host.runPoller.Start()
}

func (host WorkflowHost) Stop() {
	host.workflowConsumer.Stop()
	host.runPoller.Stop()
}
