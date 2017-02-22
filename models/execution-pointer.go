package models

import (
	"time"
)

// ExecutionPointer marks an execution point for a running workflow instance
type ExecutionPointer struct {
	ID              string
	StepID          int
	Active          bool
	PersistenceData interface{}
	StartTime       time.Time
	EndTime         time.Time
	EventName       string
	EventKey        string
	EventPublished  bool
	ConcurrentFork  int
	PathTerminator  bool
}
