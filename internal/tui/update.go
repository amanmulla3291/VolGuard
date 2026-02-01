package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case lvsLoadedMsg:
		m.LogicalLVs = msg.LVs
		m.Error = msg.Err
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			if m.Screen == MainMenu {
				return m, tea.Quit
			}
			m.Screen = MainMenu
			m.Cursor = 0
			return m, nil

		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}

		case "enter":
			if m.Screen == MainMenu {
				switch m.Cursor {
				case 0:
					m.Screen = LVMScreen
					return m, loadLVsCmd(m.LVMProvider)
				case 1:
					m.Screen = BackupScreen
				case 2:
					return m, tea.Quit
				}
			}
		}
	}

	return m, nil
}
