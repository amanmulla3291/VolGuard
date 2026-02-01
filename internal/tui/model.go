package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/amanmulla3291/volguard/internal/lvm"
)

type Screen int
type LVMTab int

const (
	MainMenu Screen = iota
	LVMScreen
	BackupScreen
)

const (
	LVsTab LVMTab = iota
	VGsTab
	PVsTab
)

type Model struct {
	Screen Screen
	Cursor int

	Choices []string

	// LVM
	LVMProvider lvm.Provider
	LVMTab      LVMTab

	LogicalLVs  []lvm.LogicalVolume
	VolumeVGs   []lvm.VolumeGroup
	PhysicalPVs []lvm.PhysicalVolume

	Error error
}

func NewModel(provider lvm.Provider) Model {
	return Model{
		Screen:      MainMenu,
		LVMProvider: provider,
		LVMTab:      LVsTab,
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
