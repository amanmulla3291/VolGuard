package tui

import "github.com/amanmulla3291/volguard/internal/lvm"

type lvsLoadedMsg struct {
	LVs []lvm.LogicalVolume
	Err error
}

type vgsLoadedMsg struct {
	VGs []lvm.VolumeGroup
	Err error
}

type pvsLoadedMsg struct {
	PVs []lvm.PhysicalVolume
	Err error
}
