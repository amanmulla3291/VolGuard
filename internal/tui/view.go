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
	s := "🗄 LVM Manager   [ MOCK MODE ]\n\n"
	s += m.tabHeader()
	s += "\n"

	if m.Error != nil {
		s += fmt.Sprintf("Error: %v\n\n", m.Error)
	}

	switch m.LVMTab {
	case LVsTab:
		for _, lv := range m.LogicalLVs {
			s += fmt.Sprintf("• %s  [%s]  %s  (%s)\n",
				lv.Name, lv.VG, lv.Size, lv.FS)
		}
	case VGsTab:
		for _, vg := range m.VolumeVGs {
			s += fmt.Sprintf("• %s  %s  free: %s\n",
				vg.Name, vg.Size, vg.Free)
		}
	case PVsTab:
		for _, pv := range m.PhysicalPVs {
			s += fmt.Sprintf("• %s  %s  VG:%s\n",
				pv.Name, pv.Size, pv.VG)
		}
	}

	s += "\n← → tab switch • q back\n"
	return s
}

func (m Model) tabHeader() string {
	tabs := []string{"LVs", "VGs", "PVs"}
	out := ""

	for i, t := range tabs {
		if LVMTab(i) == m.LVMTab {
			out += fmt.Sprintf("[ %s ] ", t)
		} else {
			out += fmt.Sprintf("  %s   ", t)
		}
	}
	return out
}
