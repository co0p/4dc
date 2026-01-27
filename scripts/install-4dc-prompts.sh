#!/usr/bin/env bash
set -euo pipefail

TARGET_DIR=".github/prompts"
WORKING_DIR=".4dc"
PREFIX="4dc"
REPO="https://github.com/co0p/4dc"
BRANCH="main"

PROMPT_FILES=(
  "constitution.prompt.md"
  "increment.prompt.md"
  "implement.prompt.md"
  "promote.prompt.md"
  "reflect.prompt.md"
)

echo ""
echo ">> 4dc installer"
echo "   Working directory: $(pwd)"
echo ""
echo "This script will:"
echo "  1. Download 4dc from $REPO@$BRANCH"
echo "  2. Copy prompt files to $TARGET_DIR/ with '$PREFIX-' prefix:"
for file in "${PROMPT_FILES[@]}"; do
  echo "     - $PREFIX-$file"
done
echo "  3. Create working directory $WORKING_DIR/"
echo "  4. Add $WORKING_DIR to .gitignore"
echo ""

read -p "Proceed? [y/N] " -n 1 -r
echo ""
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
  echo ">> Aborted."
  exit 0
fi

echo ""
echo ">> Downloading 4dc ($REPO@$BRANCH) ..."

TMP_DIR="$(mktemp -d)"
trap 'rm -rf "$TMP_DIR"' EXIT

curl -fsSL "$REPO/archive/$BRANCH.tar.gz" -o "$TMP_DIR/4dc.tar.gz"

echo ">> Extracting archive ..."
tar -xzf "$TMP_DIR/4dc.tar.gz" -C "$TMP_DIR"

# The archive extracts into 4dc-<branch>/
REPO_ROOT="$TMP_DIR/4dc-$BRANCH"

mkdir -p "$TARGET_DIR"

for file in "${PROMPT_FILES[@]}"; do
  SRC="$REPO_ROOT/$file"
  if [ -f "$SRC" ]; then
    echo "   - copying $file -> $TARGET_DIR/$PREFIX-$file"
    cp "$SRC" "$TARGET_DIR/$PREFIX-$file"
  else
    echo "   - warning: $file not found in repo root; skipping"
  fi
done

echo ">> Creating working directory $WORKING_DIR ..."
mkdir -p "$WORKING_DIR"

echo ">> Adding $WORKING_DIR to .gitignore ..."
if [ -f .gitignore ]; then
  if ! grep -qxF "$WORKING_DIR" .gitignore; then
    echo "$WORKING_DIR" >> .gitignore
    echo "   - added $WORKING_DIR to .gitignore"
  else
    echo "   - $WORKING_DIR already in .gitignore"
  fi
else
  echo "$WORKING_DIR" > .gitignore
  echo "   - created .gitignore with $WORKING_DIR"
fi

echo ">> 4dc: prompt files installed into $TARGET_DIR"
echo "   Configure your LLM / Copilot Chat to use these *$PREFIX*.prompt.md files."