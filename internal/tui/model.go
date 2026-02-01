package tui

import tea "github.com/charmbracelet/bubbletea"

type Screen int

const (
	MainMenu Screen = iota
	LVMMenu
	BackupMenu
)

type Model struct {
	Screen  Screen
	Cursor  int
	Choices []string
}

func NewModel() Model {
	return Model{
		Screen: MainMenu,
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
