package lvm

import (
	"strings"

	"github.com/amanmulla3291/volguard/internal/system"
)

func EnrichLVs(
	lvs []LogicalVolume,
	blocks []system.BlockDevice,
	usages []system.FSUsage,
) []LogicalVolume {

	for i, lv := range lvs {
		devicePath := "/dev/" + lv.VG + "/" + lv.Name

		// Match lsblk
		for _, b := range blocks {
			if strings.HasSuffix(devicePath, b.Name) {
				lvs[i].FS = b.FSType
				lvs[i].Mountpoint = b.Mountpoint
			}
		}

		// Match df
		for _, u := range usages {
			if u.Device == devicePath {
				lvs[i].Used = u.Used
				lvs[i].Avail = u.Avail
				lvs[i].UsePercent = u.UsePercent
			}
		}
	}

	return lvs
}
