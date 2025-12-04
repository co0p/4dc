#!/usr/bin/env bash
set -euo pipefail

# Location of this script: .../templates
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

run_script() {
  name="$1"
  script_path="$2"

  if [ -x "$script_path" ]; then
    echo "=== Running $name ==="
    (
      cd "$(dirname "$script_path")"
      "./$(basename "$script_path")"
    )
    echo "=== Finished $name ==="
    echo
  else
    echo "!!! Skipping $name: $script_path not found or not executable" >&2
  fi
}

# Increment: generates increment prompt template (if present)
run_script "increment/generate.sh" "$ROOT_DIR/increment/generate.sh"

# Design: generates design prompt template (if present)
run_script "design/generate.sh" "$ROOT_DIR/design/generate.sh"

# Implement: generates implement prompt template (if present)
run_script "implement/generate.sh" "$ROOT_DIR/implement/generate.sh"

# Improve: generates improve prompt template (if present)
run_script "improve/generate.sh" "$ROOT_DIR/improve/generate.sh"