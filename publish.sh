#!/usr/bin/env bash
set -euo pipefail

REPO="${GITHUB_REPOSITORY:-anku0669/orbita}"
VISIBILITY="${GITHUB_VISIBILITY:-public}"

if ! command -v gh >/dev/null 2>&1; then
  echo "GitHub CLI (gh) is required. Install it from https://cli.github.com/ and run: gh auth login" >&2
  exit 1
fi
if ! gh auth status >/dev/null 2>&1; then
  echo "Authenticate first with: gh auth login" >&2
  exit 1
fi

if gh repo view "$REPO" >/dev/null 2>&1; then
  echo "Repository $REPO already exists; refusing to overwrite it." >&2
  exit 1
fi

# Initialize Git only when this folder has not already been initialized.
if [ ! -d .git ]; then
  git init -b main
fi

# Configure a local commit identity when Git has no global identity yet.
if ! git config user.name >/dev/null; then
  git config user.name "ORBITA Publisher"
fi
if ! git config user.email >/dev/null; then
  git config user.email "orbita-publisher@users.noreply.github.com"
fi

git add .
if ! git rev-parse --verify HEAD >/dev/null 2>&1; then
  git commit -m "Initial ORBITA research prototype"
elif ! git diff --cached --quiet; then
  git commit -m "Publish ORBITA research prototype"
fi

# Create the remote and push the current branch.
gh repo create "$REPO" "--$VISIBILITY" --source=. --remote=origin --push

echo "Published: https://github.com/$REPO"
echo "Enable GitHub Pages or connect a deployment host for a public web console; this release is a CLI, not a hosted service."
