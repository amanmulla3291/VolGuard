package lvm

type PhysicalVolume struct {
	Name string
	Size string
	VG   string
}

type VolumeGroup struct {
	Name string
	Size string
	Free string
}

type LogicalVolume struct {
	Name string
	VG   string
	Size string
	FS   string
}
