package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case lvsLoadedMsg:
		m.LogicalLVs = msg.LVs
		m.Error = msg.Err
		return m, nil

	case vgsLoadedMsg:
		m.VolumeVGs = msg.VGs
		m.Error = msg.Err
		return m, nil

	case pvsLoadedMsg:
		m.PhysicalPVs = msg.PVs
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

		case "left", "h":
			if m.Screen == LVMScreen && m.LVMTab > 0 {
				m.LVMTab--
				return m, m.loadCurrentTab()
			}

		case "right", "l", "tab":
			if m.Screen == LVMScreen && m.LVMTab < PVsTab {
				m.LVMTab++
				return m, m.loadCurrentTab()
			}

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
					return m, m.loadCurrentTab()
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

func (m Model) loadCurrentTab() tea.Cmd {
	switch m.LVMTab {
	case LVsTab:
		return loadLVsCmd(m.LVMProvider)
	case VGsTab:
		return loadVGsCmd(m.LVMProvider)
	case PVsTab:
		return loadPVsCmd(m.LVMProvider)
	default:
		return nil
	}
}
