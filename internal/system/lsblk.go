package system

import (
	"context"
	"encoding/json"
)

type BlockDevice struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	FSType     string `json:"fstype"`
	Mountpoint string `json:"mountpoint"`
}

type lsblkOutput struct {
	Blockdevices []BlockDevice `json:"blockdevices"`
}

func ListBlockDevices(ctx context.Context, exec Executor) ([]BlockDevice, error) {
	out, err := exec.Run(
		ctx,
		"lsblk",
		"-J",
		"-o", "NAME,TYPE,FSTYPE,MOUNTPOINT",
	)
	if err != nil {
		return nil, err
	}

	var parsed lsblkOutput
	if err := json.Unmarshal(out, &parsed); err != nil {
		return nil, err
	}

	return parsed.Blockdevices, nil
}
