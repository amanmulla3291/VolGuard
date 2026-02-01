package tui

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/amanmulla3291/volguard/internal/lvm"
)

func loadLVsCmd(provider lvm.Provider) tea.Cmd {
	return func() tea.Msg {
		lvs, err := provider.ListLVs(context.Background())
		return lvsLoadedMsg{LVs: lvs, Err: err}
	}
}

func loadVGsCmd(provider lvm.Provider) tea.Cmd {
	return func() tea.Msg {
		vgs, err := provider.ListVGs(context.Background())
		return vgsLoadedMsg{VGs: vgs, Err: err}
	}
}

func loadPVsCmd(provider lvm.Provider) tea.Cmd {
	return func() tea.Msg {
		pvs, err := provider.ListPVs(context.Background())
		return pvsLoadedMsg{PVs: pvs, Err: err}
	}
}
