package interfaces

import (
	"github.com/danielgerlag/goflow/models"
)

type StepBody interface {
	Run(context models.StepExectuionContext) models.ExecutionResult
}
