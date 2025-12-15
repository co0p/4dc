#!/usr/bin/env bash
set -euo pipefail

TARGET_DIR=".github/prompts"
PREFIX="4dc"
REPO="https://github.com/co0p/4dc"
BRANCH="main"

echo ">> 4dc: installing prompt files into $TARGET_DIR"
echo "   Working directory: $(pwd)"

TMP_DIR="$(mktemp -d)"
trap 'rm -rf "$TMP_DIR"' EXIT

echo ">> Downloading 4dc ($REPO@$BRANCH) ..."
curl -fsSL "$REPO/archive/$BRANCH.tar.gz" -o "$TMP_DIR/4dc.tar.gz"

echo ">> Extracting archive ..."
tar -xzf "$TMP_DIR/4dc.tar.gz" -C "$TMP_DIR"

# The archive extracts into 4dc-<branch>/
REPO_ROOT="$TMP_DIR/4dc-$BRANCH"

mkdir -p "$TARGET_DIR"

PROMPT_FILES=(
  "create-constitution.prompt.md"
  "increment.prompt.md"
  "design.prompt.md"
  "implement.prompt.md"
  "improve.prompt.md"
)

for file in "${PROMPT_FILES[@]}"; do
  SRC="$REPO_ROOT/$file"
  if [ -f "$SRC" ]; then
    echo "   - copying $file -> $TARGET_DIR/$PREFIX-$file"
    cp "$SRC" "$TARGET_DIR/$PREFIX-$file"
  else
    echo "   - warning: $file not found in repo root; skipping"
  fi
done

echo ">> 4dc: prompt files installed into $TARGET_DIR"
echo "   Configure your LLM / Copilot Chat to use these *$PREFIX*.prompt.md files."