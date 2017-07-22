package models

type ExecutionResult struct {
	Proceed         bool
	PersistenceData interface{}
	BranchValues    []interface{}
}
