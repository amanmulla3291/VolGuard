package lvm

import (
	"os"
	"os/exec"
)

type Capability struct {
	LVMAvailable bool
	ReadOnly     bool
	Reason       string
}

func DetectCapability() Capability {
	if os.Geteuid() != 0 {
		return Capability{
			LVMAvailable: false,
			ReadOnly:     true,
			Reason:       "not running as root",
		}
	}

	if _, err := exec.LookPath("lvs"); err != nil {
		return Capability{
			LVMAvailable: false,
			ReadOnly:     true,
			Reason:       "lvm2 not installed (lvs not found)",
		}
	}

	return Capability{
		LVMAvailable: true,
		ReadOnly:     true, // Phase 2 is strictly read-only
		Reason:       "read-only mode (discovery only)",
	}
}
