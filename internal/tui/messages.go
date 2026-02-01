package tui

import "github.com/amanmulla3291/volguard/internal/lvm"

type lvsLoadedMsg struct {
	LVs []lvm.LogicalVolume
	Err error
}
