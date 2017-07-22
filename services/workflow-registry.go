package services

import (
	"errors"

	"github.com/danielgerlag/goflow/models"
)

type workflowRegistry struct {
	defs []models.WorkflowDefinition
}

func (registry *workflowRegistry) GetDefinition(workflowID string, version int) (*models.WorkflowDefinition, error) {
	for _, v := range registry.defs {
		if v.ID == workflowID && v.Version == version {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}
