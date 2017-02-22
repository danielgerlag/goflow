package services

import (
	"github.com/danielgerlag/goflow/models"
	"github.com/nu7hatch/gouuid"
)

// MemoryPersistenceProvider perists workflows
type MemoryPersistenceProvider struct {
	instances []models.WorkflowInstance
}

func (provider MemoryPersistenceProvider) CreateNewWorkflow(workflow models.WorkflowInstance) string {
	var u, err = uuid.NewV4()
	workflow.ID = u.String()
	var slice = append(provider.instances, workflow)
	return workflow.ID
}

func (provider MemoryPersistenceProvider) PersistWorkflow(workflow models.WorkflowInstance) {
	for i, v := range provider.instances {
		if v.ID == workflow.ID {
			var a = append(provider.instances[:i], provider.instances[i+1:]...)
		}
	}
	var slice = append(provider.instances, workflow)
}

func (provider MemoryPersistenceProvider) GetRunnableInstances() []string {
	var result = make([]string, 0)
	for i, v := range provider.instances {
		if v.Status == models.Runnable {
			result = append(result, v.ID)
		}
	}
	return result
}

func (provider MemoryPersistenceProvider) GetWorkflowInstance(id string) *models.WorkflowInstance {
	for i, v := range provider.instances {
		if v.ID == id {
			return &v
		}
	}
	return nil
}
