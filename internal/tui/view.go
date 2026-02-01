package tui

import "fmt"

func (m Model) View() string {
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
