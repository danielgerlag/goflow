package models

type ExecutionResult struct {
	Proceed         bool
	OutcomeValue    interface{}
	PersistenceData interface{}
}
