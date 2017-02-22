package services

import (
	"testing"

	"github.com/danielgerlag/goflow/models"
)

func TestCreateNewWorkflow(t *testing.T) {

	var prov = &MemoryPersistenceProvider{}
	var wf = models.WorkflowInstance{
		Description: "wf1",
	}

	var id = prov.CreateNewWorkflow(wf)
	if id == "" {
		t.FailNow()
	}

	if len(prov.instances) != 1 {
		t.FailNow()
	}
}

func TestPersistWorkflow(t *testing.T) {

	var prov = &MemoryPersistenceProvider{}
	var wf = models.WorkflowInstance{
		Description: "wf1",
	}

	var id = prov.CreateNewWorkflow(wf)

	var wf2 = models.WorkflowInstance{
		ID:          id,
		Description: "wf2",
	}

	prov.PersistWorkflow(wf2)

	if len(prov.instances) != 1 {
		t.FailNow()
	}

	var wf3, _ = prov.GetWorkflowInstance(id)

	if wf3.Description != "wf2" {
		t.FailNow()
	}

}
