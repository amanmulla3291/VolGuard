package lvm

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/amanmulla3291/volguard/internal/system"
)

/* =========================
   REAL PROVIDER
========================= */

type RealProvider struct {
	Exec system.Executor
}

func NewRealProvider(exec system.Executor) *RealProvider {
	return &RealProvider{Exec: exec}
}

/* =========================
   LVS JSON
========================= */

type lvsReport struct {
	Report []struct {
		LV []struct {
			LVName string `json:"lv_name"`
			VGName string `json:"vg_name"`
			LVSize string `json:"lv_size"`
			LVAttr string `json:"lv_attr"`
		} `json:"lv"`
	} `json:"report"`
}

func (r *RealProvider) ListLVs(ctx context.Context) ([]LogicalVolume, error) {
	out, err := r.Exec.Run(
		ctx,
		"lvs",
		"--reportformat", "json",
		"--units", "g",
		"--nosuffix",
	)
	if err != nil {
		return nil, err
	}

	var report lvsReport
	if err := json.Unmarshal(out, &report); err != nil {
		return nil, errors.New("failed to parse lvs JSON output")
	}

	var lvs []LogicalVolume
	for _, rep := range report.Report {
		for _, lv := range rep.LV {
			lvs = append(lvs, LogicalVolume{
				Name: lv.LVName,
				VG:   lv.VGName,
				Size: lv.LVSize + "G",
				FS:   "unknown",
			})
		}
	}

	return lvs, nil
}

/* =========================
   VGS JSON
========================= */

type vgsReport struct {
	Report []struct {
		VG []struct {
			VGName string `json:"vg_name"`
			VGSize string `json:"vg_size"`
			VGFree string `json:"vg_free"`
		} `json:"vg"`
	} `json:"report"`
}

func (r *RealProvider) ListVGs(ctx context.Context) ([]VolumeGroup, error) {
	out, err := r.Exec.Run(
		ctx,
		"vgs",
		"--reportformat", "json",
		"--units", "g",
		"--nosuffix",
	)
	if err != nil {
		return nil, err
	}

	var report vgsReport
	if err := json.Unmarshal(out, &report); err != nil {
		return nil, errors.New("failed to parse vgs JSON output")
	}

	var vgs []VolumeGroup
	for _, rep := range report.Report {
		for _, vg := range rep.VG {
			vgs = append(vgs, VolumeGroup{
				Name: vg.VGName,
				Size: vg.VGSize + "G",
				Free: vg.VGFree + "G",
			})
		}
	}

	return vgs, nil
}

/* =========================
   PVS JSON
========================= */

type pvsReport struct {
	Report []struct {
		PV []struct {
			PVName string `json:"pv_name"`
			PVSize string `json:"pv_size"`
			VGName string `json:"vg_name"`
		} `json:"pv"`
	} `json:"report"`
}

func (r *RealProvider) ListPVs(ctx context.Context) ([]PhysicalVolume, error) {
	out, err := r.Exec.Run(
		ctx,
		"pvs",
		"--reportformat", "json",
		"--units", "g",
		"--nosuffix",
	)
	if err != nil {
		return nil, err
	}

	var report pvsReport
	if err := json.Unmarshal(out, &report); err != nil {
		return nil, errors.New("failed to parse pvs JSON output")
	}

	var pvs []PhysicalVolume
	for _, rep := range report.Report {
		for _, pv := range rep.PV {
			pvs = append(pvs, PhysicalVolume{
				Name: pv.PVName,
				Size: pv.PVSize + "G",
				VG:   pv.VGName,
			})
		}
	}

	return pvs, nil
}
