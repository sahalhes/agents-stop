# Agent instructions

## Scope

- This project provides a small, system-wide `agents stop` CLI for Windows, macOS, and Linux.
- Keep the command and its process matching explicit. Shared runtimes such as Node and Python require a recognizable agent command-line marker.
- Preserve `--dry-run`; stopping is forceful and includes descendants of matched agent processes.
- Do not add dependencies unless required for a concrete feature.

## Changes

- Keep platform-specific process discovery and termination in files with appropriate Go build tags.
- Update the Phase 1 implementation notes in `README.md` when behavior or supported agent matchers change.
- Keep later release packaging and automation work out of Phase 1 documentation beyond the phase plan.
- Use `gofmt` on Go changes. Build for Windows, macOS, and Linux before considering the cross-platform implementation complete.
