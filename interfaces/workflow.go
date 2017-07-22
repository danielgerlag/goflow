package interfaces

type Workflow interface {
	Id() string
	Version() int
	Build(builder WorkflowBuilder)
}
