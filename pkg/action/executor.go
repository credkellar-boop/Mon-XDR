package action

import (
	"context"
	"github.com/credkellar-boop/Mon-XDR/pkg/schema"
)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

// These must be capitalized to be "public"
func (e *Executor) KillProcess(pid int) error {
	// Implementation logic
	return nil
}

func (e *Executor) CollectTelemetry(data schema.Action) error {
	// Implementation logic
	return nil
}

func (e *Executor) Execute(ctx context.Context, actionData schema.Action) error {
	return nil
}
