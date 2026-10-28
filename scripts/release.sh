#!/usr/bin/env bash
set -euo pipefail

# Renders the agent templates in src/agents/ using the shared blocks in
# src/parts/, then publishes the result to dist/agents/.
#
# A template line of the form
#
#   <!-- include: language.md -->
#
# is replaced by the contents of src/parts/language.md. Include paths resolve
# inside src/parts/ only. A missing include or an unresolved marker stops the
# release.
#
# dist/agents/ is what the installer ships to users. Release writes there only.

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SRC_AGENTS="$REPO_ROOT/src/agents"
PARTS="$REPO_ROOT/src/parts"
DIST_AGENTS="$REPO_ROOT/dist/agents"
VERSION="$(date +%Y-%m-%d)"

INCLUDE_PATTERN='^[[:space:]]*<!--[[:space:]]*include:[[:space:]]*([^[:space:]]+)[[:space:]]*-->[[:space:]]*$'

if [ ! -d "$SRC_AGENTS" ]; then
  echo "Error: missing template directory: ${SRC_AGENTS#"$REPO_ROOT"/}" >&2
  exit 1
fi

TMP_DIR="$(mktemp -d)"
trap 'rm -rf "$TMP_DIR"' EXIT

add_version() {
  local output="$1"
  local template="$2"
  local tmp="$output.version"

  if ! awk -v version="$VERSION" '
    BEGIN { in_frontmatter = 0; added = 0 }
    NR == 1 && $0 == "---" { in_frontmatter = 1; print; next }
    in_frontmatter && $0 == "---" {
      if (!added) { print "version: \"" version "\""; added = 1 }
      in_frontmatter = 0
      print
      next
    }
    in_frontmatter && /^mode:[[:space:]]/ && !added {
      print
      print "version: \"" version "\""
      added = 1
      next
    }
    { print }
    END { if (!added) exit 3 }
  ' "$output" > "$tmp"; then
    rm -f "$tmp"
    echo "Error: could not add version to ${template#"$REPO_ROOT"/}" >&2
    exit 1
  fi

  mv "$tmp" "$output"
}

render() {
  local template="$1"
  local output="$2"
  local line include path

  : > "$output"

  while IFS= read -r line || [ -n "$line" ]; do
    if [[ "$line" =~ $INCLUDE_PATTERN ]]; then
      include="${BASH_REMATCH[1]}"

      case "$include" in
        /*|*..*)
          echo "Error: include must stay inside src/parts/: $include" >&2
          exit 1
          ;;
      esac

      path="$PARTS/$include"
      if [ ! -f "$path" ]; then
        echo "Error: missing include '$include' in ${template#"$REPO_ROOT"/}" >&2
        exit 1
      fi

      cat "$path" >> "$output"
      if [ -n "$(tail -c 1 "$path")" ]; then
        printf '\n' >> "$output"
      fi
    else
      printf '%s\n' "$line" >> "$output"
    fi
  done < "$template"

  if grep -q '<!--[[:space:]]*include:' "$output"; then
    echo "Error: unresolved include in ${template#"$REPO_ROOT"/}" >&2
    exit 1
  fi

  add_version "$output" "$template"
}

shopt -s nullglob
templates=("$SRC_AGENTS"/*.md)
shopt -u nullglob

if [ ${#templates[@]} -eq 0 ]; then
  echo "Error: no agent templates found in ${SRC_AGENTS#"$REPO_ROOT"/}" >&2
  exit 1
fi

echo ""
echo ">> 4dc release"
echo "   Rendering src/agents/ with src/parts/"
echo "   Version: $VERSION"
echo ""

for template in "${templates[@]}"; do
  name="$(basename "$template")"
  render "$template" "$TMP_DIR/$name"
  echo "   $name"
done

mkdir -p "$DIST_AGENTS"
cp "$TMP_DIR"/*.md "$DIST_AGENTS/"

echo ""
echo "Done. Rendered ${#templates[@]} agents into dist/agents/"
echo ""
echo "Review the diff, then commit with:"
echo "  git add src/ dist/ && git commit -m 'release: publish updated agents'"
