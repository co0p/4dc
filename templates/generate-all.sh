#!/usr/bin/env bash
set -euo pipefail

# Generate all skill files from root-level templates.
#
# Usage (from repo root):
#   ./templates/generate-all.sh
#
# Output files are written to skills/<name>/SKILL.md:
#   - skills/constitution/SKILL.md
#   - skills/increment/SKILL.md
#   - skills/plan/SKILL.md
#   - skills/implement/SKILL.md
#   - skills/promote/SKILL.md

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

# Splice root template blocks: replace {{TEMPLATE:name}} with contents of templates/name.md
splice_templates() {
  local content
  content="$(cat)"
  local marker fragment path
  while IFS= read -r marker; do
    fragment="${marker#\{\{TEMPLATE:}"
    fragment="${fragment%\}\}}"
    path="${SCRIPT_DIR}/${fragment}.md"
    if [ -f "$path" ]; then
      local replacement
      replacement="$(cat "$path")"
      content="${content//${marker}/${replacement}}"
    else
      echo "!!! Template block not found: $path" >&2
    fi
  done < <(echo "$content" | grep -o '{{TEMPLATE:[^}]*}}' | sort -u)
  printf '%s' "$content"
}

# Replace template variables
render() {
  sed -e "s/{{VERSION}}/${COMMIT_HASH}/g" \
      -e "s/{{GENERATED_AT}}/${GENERATED_AT}/g" \
      -e "s#{{SOURCE_URL}}#${SOURCE_URL}#g"
}

# Generate a skill file from template.md into skills/<name>/SKILL.md
generate_skill() {
  local name="$1"
  local template="${SCRIPT_DIR}/${name}.md"
  local skill_dir="${ROOT_DIR}/skills/${name}"
  local output="${skill_dir}/SKILL.md"

  if [ -f "$template" ]; then
    echo "Generating skills/${name}/SKILL.md..."
    mkdir -p "$skill_dir"
    splice_shared < "$template" | splice_templates | render > "$output"
    echo "  Wrote: $output"
  else
    echo "!!! Skipping $name: $template not found" >&2
  fi
}

# Generate all skills
generate_skill "constitution"
generate_skill "increment"
generate_skill "plan"
generate_skill "implement"
generate_skill "promote"

echo
echo "Done. Generated skills:"
find "${ROOT_DIR}/skills" -name 'SKILL.md' 2>/dev/null | sort || echo "No SKILL.md files found"