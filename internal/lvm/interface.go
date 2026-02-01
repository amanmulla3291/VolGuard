package lvm

import "context"

type Provider interface {
	ListPVs(ctx context.Context) ([]PhysicalVolume, error)
	ListVGs(ctx context.Context) ([]VolumeGroup, error)
	ListLVs(ctx context.Context) ([]LogicalVolume, error)
}
