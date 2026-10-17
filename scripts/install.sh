#!/usr/bin/env bash
set -euo pipefail

for cmd in curl tar mktemp; do
  if ! command -v "$cmd" >/dev/null 2>&1; then
    echo "Error: required command not found: $cmd" >&2
    exit 1
  fi
done

REPO="co0p/4dc"
BRANCH="main"
FORCE=0

if [ "${1:-}" = "--force" ]; then
  FORCE=1
fi

echo ""
echo ">> 4dc installer"
echo "   Working directory: $(pwd)"
echo ""
echo "This script will:"
echo "  1. Download 4dc from https://github.com/$REPO at branch $BRANCH"
echo "  2. Copy dist/agents/ into your project's .opencode/agents/"
echo "  3. Copy opencode.jsonc into your project (skipped if it exists; use --force to overwrite)"
echo "  4. Add .agent/ to .gitignore"
echo ""

read -r -p "Continue? [y/N] " answer </dev/tty
if [[ "$answer" != "y" && "$answer" != "Y" ]]; then
  echo "Aborted. No confirmation given."
  exit 0
fi

TMPDIR=$(mktemp -d)
trap 'rm -rf "$TMPDIR"' EXIT

echo ""
echo "Downloading 4dc..."
curl -fsSL "https://github.com/$REPO/archive/refs/heads/$BRANCH.tar.gz" \
  | tar -xz -C "$TMPDIR" --strip-components=1

echo "Installing agents..."
mkdir -p .opencode/agents
cp -r "$TMPDIR/dist/agents/." .opencode/agents/

if [ ! -f opencode.jsonc ] || [ "$FORCE" = "1" ]; then
  cp "$TMPDIR/opencode.jsonc" opencode.jsonc
  echo "Copied opencode.jsonc"
else
  echo "Skipped opencode.jsonc (already exists; use --force to overwrite)"
fi

if ! grep -qxF '.agent/' .gitignore 2>/dev/null; then
  echo '.agent/' >> .gitignore
  echo "Added .agent/ to .gitignore"
fi

echo ""
echo "Done. Tell OpenCode: 'Use the ideate agent.'"
