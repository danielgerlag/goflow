package interfaces

const (
	WorkflowQueue = iota
	EventQueue    = iota
)

type QueueProvider interface {
	QueueForProcessing(queue int, id string)
	DequeueForProcessing(queue int) (received bool, id string)
}
