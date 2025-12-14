#!/usr/bin/env bash
set -euo pipefail

# Generate the self-contained implement.prompt.md from the template fragments.
#
# Usage (from repo root):
#   ./templates/implement/generate.sh > implement.prompt.md
#
# Or:
#   ./templates/implement/generate.sh
#
# To write to stdout instead of a file, set:
#   WRITE_TO_STDOUT=1 ./templates/implement/generate.sh

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
TEMPLATES_DIR="${ROOT_DIR}/templates/implement"
OUT_FILE="${ROOT_DIR}/implement.prompt.md"

# If you want to write to stdout instead of a file, set WRITE_TO_STDOUT=1
WRITE_TO_STDOUT="${WRITE_TO_STDOUT:-0}"

generate() {
  cat "${TEMPLATES_DIR}/00-header.md"
  echo
  cat "${TEMPLATES_DIR}/00a-subject-scope.md"
  echo
  cat "${TEMPLATES_DIR}/01-persona.md"
  echo
    cat "${TEMPLATES_DIR}/02-goal.md"
  echo
  cat "${TEMPLATES_DIR}/03-inputs.md"
  echo
  cat "${TEMPLATES_DIR}/04-process.md"
  echo
  cat "${TEMPLATES_DIR}/05-output-and-examples.md"
}

if [ "$WRITE_TO_STDOUT" = "1" ]; then
  generate
else
  generate > "$OUT_FILE"
  echo "Wrote implement prompt to: $OUT_FILE"
fi