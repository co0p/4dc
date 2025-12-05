#!/usr/bin/env bash
set -euo pipefail

# Generate the self-contained improve.prompt.md from the template fragments.
#
# Usage (from repo root):
#   ./templates/improve/generate.sh > improve.prompt.md
#
# Or:
#   ./templates/improve/generate.sh
#
# To force stdout:
#   WRITE_TO_STDOUT=1 ./templates/improve/generate.sh

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
TEMPLATES_DIR="${ROOT_DIR}/templates/improve"
OUT_FILE="${ROOT_DIR}/improve.prompt.md"

WRITE_TO_STDOUT="${WRITE_TO_STDOUT:-0}"

generate() {
  cat "${TEMPLATES_DIR}/00-header.md"
  echo
  cat "${TEMPLATES_DIR}/01-persona.md"
  echo
  cat "${TEMPLATES_DIR}/02-lenses.md"
  echo
  cat "${TEMPLATES_DIR}/03-process.md"
  echo
  cat "${TEMPLATES_DIR}/04-interaction.md"
  echo
  cat "${TEMPLATES_DIR}/05-output.md"
  echo
  cat "${TEMPLATES_DIR}/06-adr.md"
  echo
}

if [ "$WRITE_TO_STDOUT" = "1" ]; then
  generate
else
  generate > "$OUT_FILE"
  echo "Wrote improve prompt to: $OUT_FILE"
fi