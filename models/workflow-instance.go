package models

import (
	"time"
)

const (
	Runnable   = iota
	Suspended  = iota
	Complete   = iota
	Terminated = iota
)

// WorkflowInstance is an actual instance of a workflow definition
type WorkflowInstance struct {
	ID                   string
	WorkflowDefinitionID string
	Version              int
	Description          string
	NextExecution        time.Time
	Status               int
	ExecutionPointers    []ExecutionPointer
	Data                 interface{}
	CreateTime           time.Time
	CompleteTime         time.Time
}
