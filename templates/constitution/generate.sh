#!/usr/bin/env bash
set -euo pipefail

# Generate the self-contained create-constitution.prompt.md from the template fragments.
#
# Usage (from repo root):
#   ./templates/constitution/generate.sh > create-constitution.prompt.md
#
# Or:
#   ./templates/constitution/generate.sh
#
# To force stdout:
#   WRITE_TO_STDOUT=1 ./templates/constitution/generate.sh

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
TEMPLATES_DIR="${ROOT_DIR}/templates/constitution"
OUT_FILE="${ROOT_DIR}/create-constitution.prompt.md"

WRITE_TO_STDOUT="${WRITE_TO_STDOUT:-0}"
COMMIT_HASH="$(git -C "${ROOT_DIR}" rev-parse --short HEAD 2>/dev/null || echo unknown)"
GENERATED_AT="$(date -u +%Y-%m-%dT%H:%M:%SZ)"
SOURCE_URL="https://github.com/co0p/4dc"

render() {
  sed -e "s/{{VERSION}}/${COMMIT_HASH}/g" \
      -e "s/{{GENERATED_AT}}/${GENERATED_AT}/g" \
      -e "s#{{SOURCE_URL}}#${SOURCE_URL}#g"
}

generate() {
  render < "${TEMPLATES_DIR}/00-header.md"
  echo
  render < "${TEMPLATES_DIR}/01-subject-scope.md"
  echo
  render < "${TEMPLATES_DIR}/02-persona.md"
  echo
  render < "${TEMPLATES_DIR}/03-goal.md"
  echo
  render < "${TEMPLATES_DIR}/04-process.md"
  echo
  render < "${TEMPLATES_DIR}/05-output.md"
  echo
}

if [ "$WRITE_TO_STDOUT" = "1" ]; then
  generate
else
  generate > "$OUT_FILE"
  echo "Wrote constitution prompt to: $OUT_FILE"
fi