# Scheduled development task

Use this file as the task prompt for a scheduled coding agent working in this repository. The scheduler supplies the cadence and runs the agent from the repository root. This file does not install or start a cron job.

## Goal for each run

Make one small, useful improvement toward the current phase in `README.md`. Read `AGENTS.md`, `README.md`, and the relevant source before choosing work. Prefer a concrete correctness issue, a meaningful test for existing behavior, or a documentation gap. If the current phase is complete, report that finding; do not start a later phase without an explicit project decision.

## Workflow

1. Check the working tree and recent commits. Preserve unrelated or uncommitted work. If a clean, isolated change is not possible, report the blocker and stop.
2. Choose one scoped task and implement it. Keep agent process matching explicit, preserve `--dry-run`, and account for descendants when stopping processes. Follow the platform build tag and documentation rules in `AGENTS.md`.
3. Run `gofmt` on changed Go files. Run `go test ./...` and build for Windows, macOS, and Linux using `GOOS=windows`, `GOOS=darwin`, and `GOOS=linux` with `go build -o` paths under a temporary directory. Do not run `agents stop` without `--dry-run` as a test.
4. Review the diff for scope and accidental files. If checks pass, commit only this run's changes with a descriptive message. Push the commit to the current branch only when the scheduled runner is configured and authorized to push. Never force push.
5. Report the change, test results, commit hash, and any remaining issue. If checks fail, fix them or report the failure without committing an unverified change.

Keep each run bounded. Do not add release packaging or automation to Phase 1 implementation notes beyond the existing phase plan. Do not add dependencies without a concrete need.
