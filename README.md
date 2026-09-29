# agents stop

`agents stop` scans the whole machine for supported AI coding agent processes and forcefully terminates them and their child processes. It runs on Windows, macOS, and Linux. `--dry-run` previews the matches.

## Phase plan

1. **Cross-platform CLI (current phase):** Go command, OS-specific process enumeration and termination, explicit matching, dry-run, core usage documentation.
2. **Release packaging:** versioned archives for supported OS/CPU combinations, checksums, install/uninstall instructions, and package-manager distribution.
3. **Release automation and trust:** automated tagged releases, reproducible build configuration, signing where supported, and published provenance.

This repository currently contains Phase 1 only. Later phases are intentionally left for follow-up work.

## Phase 1 use

Build from source with Go 1.22 or newer:

```sh
go build -o agents .
./agents stop --dry-run
./agents stop
```

On Windows, use `go build -o agents.exe .` and run `agents.exe stop --dry-run`.

Supported process names: Codex, Claude Code, OpenCode, Goose, Herd, Aider, Gemini CLI, Cline, and Kiro CLI. For Node/Python hosts, the command line must contain a known agent marker. A match includes its descendants so task subprocesses are stopped too. This is system-wide and may include other users' processes; operating system permissions can prevent termination. The command reports those failures.

## Phase 1 implementation notes

- No third-party Go dependencies.
- Windows enumeration uses built-in PowerShell CIM and termination uses `taskkill /T /F` to include descendants.
- macOS/Linux enumeration uses `ps` and reads its argument column without the PID, parent PID, or executable column for marker matching; termination sends `SIGKILL`.
- `--dry-run` is recommended before the first stop on a machine.
- If any matched process cannot be stopped, the command reports each failure and exits with a nonzero status.
- Agent matching is maintained in `main.go`; direct executable names match explicitly, while Node/Python processes require one of the listed agent command markers. Keep matching narrow and update this document when supported matchers change.
