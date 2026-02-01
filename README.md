# VolGuard

VolGuard is a production-grade Terminal User Interface (TUI) tool for safely managing
Linux Logical Volume Manager (LVM) volumes and performing snapshot-based backups.

> ⚠️ **Important:** This project is under active development.
> Destructive operations are not implemented yet.

## Features (Phase 1)

- Clean Bubble Tea–based TUI
- LVM explorer (LV / VG / PV views)
- Mock LVM provider for safe development
- Loading indicators and status bar
- Environment and privilege awareness
- Designed for enterprise-grade safety

## Running (Development)

```bash
go run ./cmd
```
Modes

MOCK MODE (default):

Used in GitHub Codespaces and non-LVM environments

No system commands executed

REAL MODE (Phase 2+):

Requires root

Uses LVM JSON output

Snapshot-only, fail-safe operations

Roadmap

Phase 2: Real LVM discovery and snapshot support

Phase 3: Backup & restore workflows

Phase 4: Hardening and production readiness

Disclaimer

VolGuard is designed to fail closed.
No destructive action will run without explicit confirmation.