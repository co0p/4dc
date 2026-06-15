#!/usr/bin/env bash
set -euo pipefail

TARGET_SKILLS_DIR=".github/skills"
ORCHESTRATOR=".github/AGENTS.md"
WORKING_DIR=".agent"
REPO="https://github.com/co0p/4dc"
BRANCH="main"

SKILL_NAMES=(
  "constitution"
  "increment"
  "plan"
  "implement"
  "promote"
)

echo ""
echo ">> 4dc installer"
echo "   Working directory: $(pwd)"
echo ""
echo "This script will:"
echo "  1. Download 4dc from $REPO@$BRANCH"
echo "  2. Copy skill files to $TARGET_SKILLS_DIR/<name>/SKILL.md:"
for name in "${SKILL_NAMES[@]}"; do
  echo "     - $TARGET_SKILLS_DIR/$name/SKILL.md"
done
echo "  3. Copy orchestrator to $ORCHESTRATOR"
echo "  Note: Both skills and orchestrator are installed under .github/ (hidden directory)"
echo "  4. Create working directory $WORKING_DIR/"
echo "  5. Add $WORKING_DIR to .gitignore"
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

# Install skills
for name in "${SKILL_NAMES[@]}"; do
  SRC="$REPO_ROOT/skills/$name/SKILL.md"
  DEST_DIR="$TARGET_SKILLS_DIR/$name"
  DEST="$DEST_DIR/SKILL.md"
  if [ -f "$SRC" ]; then
    mkdir -p "$DEST_DIR"
    echo "   - copying skills/$name/SKILL.md -> $DEST"
    cp "$SRC" "$DEST"
  else
    echo "   - warning: skills/$name/SKILL.md not found in repo; skipping"
  fi
done

# Install orchestrator
ORCH_SRC="$REPO_ROOT/AGENTS.md"
mkdir -p ".github"
if [ -f "$ORCH_SRC" ]; then
  echo "   - copying AGENTS.md -> $ORCHESTRATOR"
  cp "$ORCH_SRC" "$ORCHESTRATOR"
else
  echo "   - warning: AGENTS.md not found in repo; skipping"
fi

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

echo ""
echo ">> 4dc installed."
echo "   Orchestrator: $ORCHESTRATOR"
echo "   Skills:       $TARGET_SKILLS_DIR/"
echo "   Working dir:  $WORKING_DIR/ (gitignored)"
echo ""
echo "   Your agent reads .github/AGENTS.md to detect the current phase and loads the"
echo "   matching skill from .github/skills/<phase>/SKILL.md automatically."
