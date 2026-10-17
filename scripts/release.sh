#!/usr/bin/env bash
set -euo pipefail

# Copies the working agent files from src/agents/ into dist/agents/ and
# .opencode/agents/. Run this when the agents are ready to publish.
#
# dist/agents/ is what the installer ships to users.
# .opencode/agents/ is what OpenCode loads locally so the agents can
# dog-food themselves in this repository.

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

echo ""
echo ">> 4dc release"
echo "   Copying src/agents/ -> dist/agents/"
echo "   Copying src/agents/ -> .opencode/agents/"
echo ""

mkdir -p "$REPO_ROOT/dist/agents" "$REPO_ROOT/.opencode/agents"

cp "$REPO_ROOT/src/agents/"*.md "$REPO_ROOT/dist/agents/"
cp "$REPO_ROOT/src/agents/"*.md "$REPO_ROOT/.opencode/agents/"

echo "Done."
echo ""
echo "Review the diff, then commit with:"
echo "  git add dist/ .opencode/agents/ && git commit -m 'release: publish updated agents'"
