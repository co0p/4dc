#!/usr/bin/env bash
set -euo pipefail

# Generate all prompt files from template.md in each folder.
#
# Usage (from repo root):
#   ./templates/generate-all.sh
#
# Output files are written to the repo root:
#   - constitution.prompt.md
#   - increment.prompt.md
#   - design.prompt.md
#   - implement.prompt.md
#   - promote.prompt.md

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"

COMMIT_HASH="$(git -C "${ROOT_DIR}" rev-parse --short HEAD 2>/dev/null || echo unknown)"
GENERATED_AT="$(date -u +%Y-%m-%dT%H:%M:%SZ)"
SOURCE_URL="https://github.com/co0p/4dc"

# Splice shared fragments: replace {{SHARED:name}} with contents of templates/shared/name.md
splice_shared() {
  local content
  content="$(cat)"
  local marker fragment path
  while IFS= read -r marker; do
    fragment="${marker#\{\{SHARED:}"
    fragment="${fragment%\}\}}"
    path="${SCRIPT_DIR}/shared/${fragment}.md"
    if [ -f "$path" ]; then
      local replacement
      replacement="$(cat "$path")"
      # Replace the marker line with the fragment contents
      content="${content//${marker}/${replacement}}"
    else
      echo "!!! Shared fragment not found: $path" >&2
    fi
  done < <(echo "$content" | grep -o '{{SHARED:[^}]*}}' | sort -u)
  printf '%s' "$content"
}

# Replace template variables
render() {
  sed -e "s/{{VERSION}}/${COMMIT_HASH}/g" \
      -e "s/{{GENERATED_AT}}/${GENERATED_AT}/g" \
      -e "s#{{SOURCE_URL}}#${SOURCE_URL}#g"
}

# Generate a prompt from template.md
generate_prompt() {
  local name="$1"
  local template="${SCRIPT_DIR}/${name}/template.md"
  local output="${ROOT_DIR}/${name}.prompt.md"

  if [ -f "$template" ]; then
    echo "Generating ${name}.prompt.md..."
    splice_shared < "$template" | render > "$output"
    echo "  Wrote: $output"
  else
    echo "!!! Skipping $name: $template not found" >&2
  fi
}

# Generate all prompts
generate_prompt "constitution"
generate_prompt "increment"
generate_prompt "implement"
generate_prompt "design"
generate_prompt "promote"

echo
echo "Done. Generated prompts:"
ls -la "${ROOT_DIR}"/*.prompt.md 2>/dev/null || echo "No .prompt.md files found"