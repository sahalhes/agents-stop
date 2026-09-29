#!/bin/bash

set -u

REPO="/root/projects/agents-stop"
LOCK="/tmp/agents-stop-codex.lock"
LOG="/root/projects/agents-stop/codex-agent.log"

# Prevent overlapping Codex runs.
exec 9>"$LOCK"

if ! flock -n 9; then
    echo "$(date -Is) - Previous Codex run still active. Skipping." >> "$LOG"
    exit 0
fi

cd "$REPO" || exit 1

echo "" >> "$LOG"
echo "========================================" >> "$LOG"
echo "$(date -Is) - Starting Codex run" >> "$LOG"

# Always work from latest dev.
git fetch origin >> "$LOG" 2>&1
git checkout dev >> "$LOG" 2>&1

# Don't start if something unexpected is already modified.
if ! git diff --quiet || ! git diff --cached --quiet; then
    echo "$(date -Is) - Working tree has existing changes. Skipping." >> "$LOG"
    exit 1
fi

git pull --ff-only origin dev >> "$LOG" 2>&1 || exit 1

codex \
  -a never \
  --sandbox workspace-write \
  -C "$REPO" \
  exec \
  "Read AGENTS.md and AUTODEVELOP.md. Continue development of this project. Inspect the current code, documentation, tests, TODOs and recent git history. Choose ONE useful unfinished improvement. Implement it and run appropriate tests. Work only on dev. Never modify or push to main. When the work is complete and tests pass, stage the relevant changes, commit them with a descriptive commit message, and push to origin dev. Do not force push. If there is no appropriate improvement or you are blocked, make no destructive changes and explain the blocker." \
  >> "$LOG" 2>&1

EXIT_CODE=$?

echo "$(date -Is) - Codex finished with exit code $EXIT_CODE" >> "$LOG"

exit "$EXIT_CODE"