package lvm

import "context"

type MockProvider struct{}

func (m *MockProvider) ListPVs(ctx context.Context) ([]PhysicalVolume, error) {
	return []PhysicalVolume{
		{Name: "/dev/sda2", Size: "100G", VG: "vg0"},
	}, nil
}

func (m *MockProvider) ListVGs(ctx context.Context) ([]VolumeGroup, error) {
	return []VolumeGroup{
		{Name: "vg0", Size: "100G", Free: "40G"},
	}, nil
}

func (m *MockProvider) ListLVs(ctx context.Context) ([]LogicalVolume, error) {
	return []LogicalVolume{
		{Name: "root", VG: "vg0", Size: "50G", FS: "ext4"},
		{Name: "home", VG: "vg0", Size: "10G", FS: "xfs"},
	}, nil
}
