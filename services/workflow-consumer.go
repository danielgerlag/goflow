package services

import (
	"time"

	"github.com/danielgerlag/goflow/interfaces"
)

type WorkflowConsumer struct {
	persistence   interfaces.PersistenceProvider
	executor      interfaces.WorkflowExecutor
	queueProvider interfaces.QueueProvider
	started       bool
}

func MakeWorkflowConsumer(persistence interfaces.PersistenceProvider, executor interfaces.WorkflowExecutor, queueProvider interfaces.QueueProvider) WorkflowConsumer {
	return WorkflowConsumer{
		persistence:   persistence,
		executor:      executor,
		queueProvider: queueProvider,
	}
}

func (consumer *WorkflowConsumer) Start() {
	consumer.started = true
	go consumer.run()
}

func (consumer *WorkflowConsumer) Stop() {
	consumer.started = false
}

func (consumer WorkflowConsumer) run() {
	for consumer.started {
		recv, id := consumer.queueProvider.DequeueForProcessing(interfaces.WorkflowQueue)

		if recv {
			go consumer.executeItem(id)

		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (consumer WorkflowConsumer) executeItem(id string) {
	workflow := consumer.persistence.GetWorkflowInstance(id)
	consumer.executor.Execute(&workflow)
	consumer.persistence.PersistWorkflow(workflow)
}
