# agents-stop

A small, system-wide `agents stop` CLI for Windows, macOS, and Linux. It finds supported AI coding agent processes and forcefully stops them, including their descendants. Use `--dry-run` to preview the processes that would be stopped.

## Phase plan

### Phase 1: Core process stopping

- Provide `agents stop`, help, version output, and the `--dry-run` option.
- Discover processes using platform-specific implementations for Windows and Unix-like systems.
- Match supported agent executables explicitly. For shared runtimes such as Node and Python, require a recognizable agent command-line marker.
- Include descendants of matched agent processes and stop descendants before their parents.
- Keep the implementation dependency-free and document supported behavior here.

**Status:** Core command and platform process implementations are present. Cross-platform builds and behavior checks remain to be completed.

### Phase 2: Reliability and compatibility

- Improve process-list parsing and error reporting for platform-specific edge cases.
- Review process matching against supported agent command lines to reduce accidental matches and missed agents.
- Verify dry-run output and descendant handling on Windows, macOS, and Linux.

### Phase 3: Distribution

- Define release artifacts and installation guidance for supported operating systems and architectures.
- Add versioned release packaging once the distribution targets are settled.

### Phase 4: Release automation

- Automate cross-platform builds and checks in CI.
- Automate publishing versioned artifacts and release notes.

## Usage

```text
agents stop [--dry-run]
agents --version
```

<<<<<<< HEAD
Stopping is forceful. Review the `--dry-run` output before stopping processes when you are unsure which agent processes are running.
=======
On Windows, use `go build -o agents.exe .` and run `agents.exe stop --dry-run`.

Supported process names: Codex, Claude Code, OpenCode, Goose, Herd, Aider, Gemini CLI, Cline, and Kiro CLI. For Node/Python hosts, the command line must contain a known agent marker. A match includes its descendants so task subprocesses are stopped too. This is system-wide and may include other users' processes; operating system permissions can prevent termination. The command reports those failures.

## Phase 1 implementation notes

- No third-party Go dependencies.
- Windows enumeration uses built-in PowerShell CIM and termination uses `taskkill /T /F` to include descendants.
- macOS/Linux enumeration uses `ps` and reads its argument column without the PID, parent PID, or executable column for marker matching; termination sends `SIGKILL`.
- `--dry-run` is recommended before the first stop on a machine.
- If any matched process cannot be stopped, the command reports each failure and exits with a nonzero status.
- Agent matching is maintained in `main.go`; direct executable names match explicitly, while Node/Python processes require one of the listed agent command markers. Keep matching narrow and update this document when supported matchers change.
>>>>>>> 31c9779f4152716d37b767dfa086c55708afe5f7
