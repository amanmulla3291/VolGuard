package tui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/amanmulla3291/volguard/internal/lvm"
)

type Screen int

const (
	MainMenu Screen = iota
	LVMScreen
	BackupScreen
)

type Model struct {
	Screen Screen
	Cursor int

	// Main menu
	Choices []string

	// LVM
	LVMProvider lvm.Provider
	LogicalLVs  []lvm.LogicalVolume
	Error       error
}

func NewModel(provider lvm.Provider) Model {
	return Model{
		Screen:      MainMenu,
		LVMProvider: provider,
		Choices: []string{
			"LVM Manager",
			"Backup & Restore",
			"Quit",
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
