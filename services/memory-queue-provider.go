package services

import (
	"time"

	"github.com/danielgerlag/goflow/interfaces"
)

var workflowChannel = make(chan string, 100)
var eventChannel = make(chan string, 100)

// MemoryQueueProvider provides internal queueing
type memoryQueueProvider struct {
}

func NewMemoryQueueProvider() interfaces.QueueProvider {
	result := new(memoryQueueProvider)
	return result
}

func (provider *memoryQueueProvider) QueueForProcessing(queue int, id string) {

	var channel chan string

	switch queue {
	case interfaces.WorkflowQueue:
		channel = workflowChannel
	case interfaces.EventQueue:
		channel = eventChannel
	}

	go func() {
		channel <- id
	}()

}

func (provider *memoryQueueProvider) DequeueForProcessing(queue int) (received bool, id string) {

	var channel chan string

	switch queue {
	case interfaces.WorkflowQueue:
		channel = workflowChannel
	case interfaces.EventQueue:
		channel = eventChannel
	}

	select {
	case res := <-channel:
		id = res
		received = true
	case <-time.After(time.Second * 1):
		received = false
	}

	return
}
