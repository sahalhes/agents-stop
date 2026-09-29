# Autonomous Development Instructions

You are continuously developing this project.

## Primary objective

Improve this project toward a stable, useful production-ready application.

Read the existing repository, README, issues, TODOs and code before deciding what to work on.

## Branch policy

Never modify or push to main.

Work only on dev or agent/* branches.

## Every run

1. Inspect the current repository state.
2. Read AGENTS.md and relevant documentation.
3. Inspect recent git history.
4. Determine the highest-value unfinished task.
5. Implement ONE coherent unit of work.
6. Run relevant tests, linting and build checks.
7. Fix problems caused by your changes.
8. Commit the completed work.
9. Push the branch.

Do not repeatedly rewrite already-working code.

## Safety

Do not:
- push to main
- delete production data
- expose secrets
- commit .env files or credentials
- modify production infrastructure unless explicitly instructed
- force push
- bypass failing tests merely to obtain a green result

Prefer small, reviewable commits.

If a task is too large, break it into smaller tasks and complete one useful piece during this run.

If blocked, document the blocker rather than making destructive assumptions.