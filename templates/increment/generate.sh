#!/usr/bin/env bash
set -euo pipefail

# Generate a self-contained increment.prompt.md at the repo root
# by concatenating the increment template parts in a defined order.

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
TEMPLATES_DIR="$ROOT_DIR/templates/increment"
OUT_FILE="$ROOT_DIR/increment.prompt.md"

echo "Generating $OUT_FILE from templates in $TEMPLATES_DIR"

PARTS=(
  "00-header.md"   # optional: YAML frontmatter for IDEs
  "01-persona.md"
  "02-inputs.md"
  "03-goal.md"
  "04-task.md"
  "05-process.md"
  "06-output-structure.md"
)

: > "$OUT_FILE"  # truncate output file

for part in "${PARTS[@]}"; do
  PART_PATH="$TEMPLATES_DIR/$part"
  if [[ -f "$PART_PATH" ]]; then
    echo "Appending $part"
    cat "$PART_PATH" >> "$OUT_FILE"
    echo -e "\n" >> "$OUT_FILE"
  else
    echo "Skipping missing $part"
  fi
done

echo "Done. Generated $OUT_FILE"