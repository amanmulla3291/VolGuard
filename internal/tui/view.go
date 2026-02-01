package tui

import "fmt"

func (m Model) View() string {

	switch m.Screen {

	case LVMScreen:
		return m.lvmView()

	default:
		return m.mainMenuView()
	}
}

func (m Model) mainMenuView() string {
	s := "🛡 VolGuard — LVM & Backup Manager\n\n"

	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = "➜"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\n↑ ↓ enter • q quit\n"
	return s
}

func (m Model) lvmView() string {
	s := "🗄 LVM Manager (Logical Volumes)\n\n"

	if m.Error != nil {
		s += fmt.Sprintf("Error: %v\n\n", m.Error)
	}

	if len(m.LogicalLVs) == 0 {
		s += "No logical volumes found.\n"
	} else {
		for _, lv := range m.LogicalLVs {
			s += fmt.Sprintf(
				"• %s  [%s]  %s  (%s)\n",
				lv.Name,
				lv.VG,
				lv.Size,
				lv.FS,
			)
		}
	}

	s += "\nq = back\n"
	return s
}
