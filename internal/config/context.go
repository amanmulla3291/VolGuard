package config

import "os"

type AppContext struct {
	AppName   string
	Version   string
	MockMode  bool
	IsRoot    bool
	Env       string

	CapabilityReason string
	ReadOnly         bool
}


func DetectContext() AppContext {
	ctx := AppContext{
		AppName: "VolGuard",
		Version: "0.1.0",
		MockMode: true,
		Env: "unknown",
	}

	if os.Geteuid() == 0 {
		ctx.IsRoot = true
	}

	if os.Getenv("CODESPACES") == "true" {
		ctx.Env = "codespaces"
	} else {
		ctx.Env = "linux"
	}

	return ctx
}
