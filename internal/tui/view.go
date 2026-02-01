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

/* =========================
   MAIN MENU VIEW
========================= */

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
	s += m.statusBar()
	return s
}

/* =========================
   LVM VIEW
========================= */

func (m Model) lvmView() string {
	s := "🗄 LVM Manager"

	if m.Context.MockMode {
		s += "   [ MOCK MODE ]"
	}

	s += "\n\n"
	s += m.tabHeader()
	s += "\n\n"

	if m.IsLoading {
		s += fmt.Sprintf("%s Loading LVM data...\n\n", m.Spinner.View())
	}

	if m.Error != nil {
		s += fmt.Sprintf("Error: %v\n\n", m.Error)
	}

	switch m.LVMTab {

	case LVsTab:
		if len(m.LogicalLVs) == 0 && !m.IsLoading {
			s += "No logical volumes found.\n"
		}
		for _, lv := range m.LogicalLVs {
			mp := lv.Mountpoint
			if mp == "" {
				mp = "-"
			}
		
			fs := lv.FS
			if fs == "" {
				fs = "unknown"
			}
		
			usage := lv.UsePercent
			if usage == "" {
				usage = "-"
			}

	s += fmt.Sprintf(
		"• %-12s VG:%-6s %-5s %-8s %-4s %s\n",
		lv.Name,
		lv.VG,
		lv.Size,
		fs,
		usage,
		mp,
	)
}


	case VGsTab:
		if len(m.VolumeVGs) == 0 && !m.IsLoading {
			s += "No volume groups found.\n"
		}
		for _, vg := range m.VolumeVGs {
			s += fmt.Sprintf(
				"• %-10s  size:%-6s  free:%-6s\n",
				vg.Name,
				vg.Size,
				vg.Free,
			)
		}

	case PVsTab:
		if len(m.PhysicalPVs) == 0 && !m.IsLoading {
			s += "No physical volumes found.\n"
		}
		for _, pv := range m.PhysicalPVs {
			s += fmt.Sprintf(
				"• %-15s  %-6s  VG:%s\n",
				pv.Name,
				pv.Size,
				pv.VG,
			)
		}
	}

	s += "\n← → tab switch • q back\n"
	s += m.statusBar()
	return s
}

/* =========================
   TAB HEADER
========================= */

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

/* =========================
   STATUS BAR
========================= */

func (m Model) statusBar() string {
	mode := "REAL"
if m.Context.MockMode {
	mode = "MOCK"
} else if m.Context.ReadOnly {
	mode = "REAL-RO"
}

	user := "user"
	if m.Context.IsRoot {
		user = "root"
	}

	return fmt.Sprintf(
		"\n[%s] env:%s user:%s v%s\n",
		mode,
		m.Context.Env,
		user,
		m.Context.Version,
	)
}

func (m Model) capabilityWarning() string {
	if m.Context.CapabilityReason == "" {
		return ""
	}

	return fmt.Sprintf(
		"⚠ %s\n\n",
		m.Context.CapabilityReason,
	)
}
