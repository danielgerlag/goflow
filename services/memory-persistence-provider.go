package services

import (
	"errors"

	linq "gopkg.in/ahmetb/go-linq.v3"

	"github.com/danielgerlag/goflow/models"
	"github.com/nu7hatch/gouuid"
)

// MemoryPersistenceProvider perists workflows
type MemoryPersistenceProvider struct {
	instances []models.WorkflowInstance
}

func (provider *MemoryPersistenceProvider) CreateNewWorkflow(workflow models.WorkflowInstance) string {
	var u, _ = uuid.NewV4()
	workflow.ID = u.String()
	provider.instances = append(provider.instances, workflow)
	return workflow.ID
}

func (provider *MemoryPersistenceProvider) PersistWorkflow(workflow models.WorkflowInstance) {
	for i, v := range provider.instances {
		if v.ID == workflow.ID {
			provider.instances = append(provider.instances[:i], provider.instances[i+1:]...)
		}
	}
	provider.instances = append(provider.instances, workflow)
}

func (provider MemoryPersistenceProvider) GetRunnableInstances() []string {

	var result []string

	linq.From(provider.instances).WhereT(func(x models.WorkflowInstance) bool {
		return x.Status == models.Runnable
	}).SelectT(func(x models.WorkflowInstance) string {
		return x.ID
	}).ToSlice(&result)

	return result
}

func (provider MemoryPersistenceProvider) GetWorkflowInstance(id string) (*models.WorkflowInstance, error) {
	for _, v := range provider.instances {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}
