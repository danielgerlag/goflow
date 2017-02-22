package services

import (
	"errors"

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
	var result = make([]string, 0)
	for _, v := range provider.instances {
		if v.Status == models.Runnable {
			result = append(result, v.ID)
		}
	}
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
