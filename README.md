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

Stopping is forceful. Review the `--dry-run` output before stopping processes when you are unsure which agent processes are running.
