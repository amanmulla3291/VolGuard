package system

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"time"
)

type Executor interface {
	Run(ctx context.Context, cmd string, args ...string) ([]byte, error)
}

type SafeExecutor struct {
	DryRun  bool
	Timeout time.Duration
}

func (e *SafeExecutor) Run(ctx context.Context, cmd string, args ...string) ([]byte, error) {
	if e.DryRun {
		return []byte("[dry-run] " + cmd), nil
	}

	if e.Timeout == 0 {
		e.Timeout = 5 * time.Second
	}

	ctx, cancel := context.WithTimeout(ctx, e.Timeout)
	defer cancel()

	c := exec.CommandContext(ctx, cmd, args...)
	var stdout, stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr

	if err := c.Run(); err != nil {
		return nil, errors.New(stderr.String())
	}

	return stdout.Bytes(), nil
}
