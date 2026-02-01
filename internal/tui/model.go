package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/amanmulla3291/volguard/internal/config"
	"github.com/amanmulla3291/volguard/internal/lvm"
)

/* =========================
   SCREENS & TABS
========================= */

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

/* =========================
   MODEL
========================= */

type Model struct {
	// UI state
	Screen Screen
	Cursor int

	Choices []string

	// App context
	Context   config.AppContext
	IsLoading bool

	// Spinner
	Spinner spinner.Model

	// LVM
	LVMProvider lvm.Provider
	LVMTab      LVMTab

	LogicalLVs  []lvm.LogicalVolume
	VolumeVGs   []lvm.VolumeGroup
	PhysicalPVs []lvm.PhysicalVolume

	Error error
}

/* =========================
   CONSTRUCTOR
========================= */

func NewModel(ctx config.AppContext, provider lvm.Provider) Model {
	sp := spinner.New()
	sp.Spinner = spinner.Dot

	return Model{
		Screen:      MainMenu,
		Context:     ctx,
		LVMProvider: provider,
		LVMTab:      LVsTab,
		Spinner:     sp,
		Choices: []string{
			"LVM Manager",
			"Backup & Restore",
			"Quit",
		},
	}
}

/* =========================
   INIT
========================= */

func (m Model) Init() tea.Cmd {
	return nil
}
