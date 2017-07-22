package services

import (
	"time"

	"github.com/danielgerlag/goflow/interfaces"
)

type RunPoller struct {
	persistence   interfaces.PersistenceProvider
	queueProvider interfaces.QueueProvider
	timer         *time.Timer
}

func (poller *RunPoller) Start() {
	poller.timer = time.AfterFunc(1*time.Second, poller.run)
}

func (poller *RunPoller) Stop() {
	poller.timer.Stop()
}

func (poller RunPoller) run() {
	workflows := poller.persistence.GetRunnableInstances()
	for _, id := range workflows {
		poller.queueProvider.QueueForProcessing(interfaces.WorkflowQueue, id)
	}
	poller.timer = time.AfterFunc(10*time.Second, poller.run)
}
