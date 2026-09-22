#!/usr/bin/env bash
set -euo pipefail

# Generate all skill files from root-level templates.
#
# Usage (from repo root):
#   ./scripts/generate-4dc.sh
#
# Output files are written to skills/<name>/SKILL.md:
#   - skills/constitution/SKILL.md
#   - skills/increment/SKILL.md
#   - skills/prototype/SKILL.md
#   - skills/plan/SKILL.md
#   - skills/implement/SKILL.md
#   - skills/subtask-plan/SKILL.md
#   - skills/adr/SKILL.md
#   - skills/tidy/SKILL.md
#   - skills/tdd-red/SKILL.md
#   - skills/tdd-green/SKILL.md
#   - skills/refactor/SKILL.md
#   - skills/promote/SKILL.md
#
# ─────────────────────────────────────────────────────────────────────────────
# IMPORTANT: Do NOT edit skills/<name>/SKILL.md directly.
#
# Those files are generated outputs. Any direct edits will be overwritten the
# next time this script runs.
#
# To change a skill:
#   1. Edit the corresponding source in templates/<name>.md
#   2. For shared content used across multiple skills, edit templates/shared/<fragment>.md
#   3. Run this script to regenerate: ./scripts/generate-4dc.sh
#   4. Review the diff in skills/ and commit both the template change and the
#      regenerated skill together.
#
# File map:
#   templates/constitution.md       → skills/constitution/SKILL.md
#   templates/increment.md          → skills/increment/SKILL.md
#   templates/implement.md          → skills/implement/SKILL.md
#   templates/subtask-plan.md       → skills/subtask-plan/SKILL.md
#   templates/prototype.md          → skills/prototype/SKILL.md
#   templates/plan.md               → skills/plan/SKILL.md
#   templates/adr.md                → skills/adr/SKILL.md
#   templates/tidy.md               → skills/tidy/SKILL.md
#   templates/tdd-red.md            → skills/tdd-red/SKILL.md
#   templates/tdd-green.md          → skills/tdd-green/SKILL.md
#   templates/refactor.md           → skills/refactor/SKILL.md
#   templates/promote.md            → skills/promote/SKILL.md
#   templates/shared/execution-contract.md  → spliced into every skill via {{SHARED:execution-contract}}
#   templates/language.md           → spliced through {{TEMPLATE:language}} for consistent language rules
#   templates/foundations/<id>.md    → spliced through {{FOUNDATION:<id>}} into selected skills
# ─────────────────────────────────────────────────────────────────────────────

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"
TEMPLATE_DIR="${ROOT_DIR}/templates"

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
    path="${TEMPLATE_DIR}/shared/${fragment}.md"
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
    path="${TEMPLATE_DIR}/${fragment}.md"
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

# Splice foundation fragments: replace {{FOUNDATION:id}} with templates/foundations/id.md
splice_foundations() {
  local content
  content="$(cat)"
  local marker foundation path replacement
  while IFS= read -r marker; do
    [ -n "$marker" ] || continue
    foundation="${marker#\{\{FOUNDATION:}"
    foundation="${foundation%\}\}}"
    path="${TEMPLATE_DIR}/foundations/${foundation}.md"
    if [ ! -f "$path" ]; then
      echo "Error: foundation fragment not found: $path" >&2
      return 1
    fi
    replacement="$(cat "$path")"
    content="${content//${marker}/${replacement}}"
  done < <(printf '%s' "$content" | grep -o '{{FOUNDATION:[^}]*}}' | sort -u || true)
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
  local template="${TEMPLATE_DIR}/${name}.md"
  local skill_dir="${ROOT_DIR}/skills/${name}"
  local output="${skill_dir}/SKILL.md"

  if [ -f "$template" ]; then
    echo "Generating skills/${name}/SKILL.md..."
    mkdir -p "$skill_dir"
    duplicate_foundation="$(grep -o '{{FOUNDATION:[^}]*}}' "$template" | sort | uniq -d || true)"
    if [ -n "$duplicate_foundation" ]; then
      echo "Error: duplicate foundation marker in $template: $duplicate_foundation" >&2
      exit 1
    fi
    splice_shared < "$template" | splice_templates | splice_shared | splice_templates | splice_foundations | render > "$output"
    if grep -Eq '\{\{(SHARED|TEMPLATE|FOUNDATION):[^}]+\}\}' "$output"; then
      echo "Error: unresolved template marker in $output" >&2
      exit 1
    fi
    if [ "$(grep -c '^## Language and Interaction Rules$' "$output")" -ne 1 ]; then
      echo "Error: expected exactly one language rules section in $output" >&2
      exit 1
    fi
    echo "  Wrote: $output"
  else
    echo "!!! Skipping $name: $template not found" >&2
  fi
}

# Generate all skills
generate_skill "constitution"
generate_skill "increment"
generate_skill "prototype"
generate_skill "plan"
generate_skill "implement"
generate_skill "subtask-plan"
generate_skill "adr"
generate_skill "tidy"
generate_skill "tdd-red"
generate_skill "tdd-green"
generate_skill "refactor"
generate_skill "promote"

echo
echo "Done. Generated skills:"
find "${ROOT_DIR}/skills" -name 'SKILL.md' 2>/dev/null | sort || echo "No SKILL.md files found"
