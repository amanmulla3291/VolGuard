package main

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/amanmulla3291/volguard/internal/config"
	"github.com/amanmulla3291/volguard/internal/lvm"
	"github.com/amanmulla3291/volguard/internal/system"
	"github.com/amanmulla3291/volguard/internal/tui"
)

func main() {
	// Detect application context (env, root, version, etc.)
	ctx := config.DetectContext()

	// Detect LVM capability
	cap := lvm.DetectCapability()

	// Select provider (fail-closed)
	var provider lvm.Provider

	if cap.LVMAvailable {
		exec := &system.SafeExecutor{
			DryRun: false,
		}

		provider = lvm.NewRealProvider(exec)

		ctx.MockMode = false
		ctx.ReadOnly = cap.ReadOnly
		ctx.CapabilityReason = cap.Reason
	} else {
		provider = &lvm.MockProvider{}

		ctx.MockMode = true
		ctx.ReadOnly = true
		ctx.CapabilityReason = cap.Reason
	}

	// Start TUI
	p := tea.NewProgram(
		tui.NewModel(ctx, provider),
	)

	_ = p.Start()
}
