package main

import (
	"fmt"

	"github.com/danielgerlag/goflow/interfaces"
	"github.com/danielgerlag/goflow/services"
)

func main() {
	fmt.Println("start")

	var prov = services.NewMemoryQueueProvider()
	var prov2 = services.NewMemoryQueueProvider()

	prov.QueueForProcessing(interfaces.WorkflowQueue, "test")
	prov.QueueForProcessing(interfaces.WorkflowQueue, "test2")
	dequeue(prov)
	dequeue(prov2)

}

func dequeue(prov interfaces.QueueProvider) {
	var recv, result = prov.DequeueForProcessing(interfaces.WorkflowQueue)

	if recv {
		fmt.Println("recv")
		fmt.Println(result)
	} else {
		fmt.Println("not recv")
	}
}
