#!/bin/bash
set -e

CURRENT_VERSION=$(cat VERSION)
echo "Current version: $CURRENT_VERSION"

IFS='.' read -r -a VERSION_PARTS <<< "$CURRENT_VERSION"
MAJOR="${VERSION_PARTS[0]}"
MINOR="${VERSION_PARTS[1]}"
PATCH="${VERSION_PARTS[2]}"

BUMP_TYPE="${1:-patch}"

case $BUMP_TYPE in
  major)
    MAJOR=$((MAJOR + 1))
    MINOR=0
    PATCH=0
    ;;
  minor)
    MINOR=$((MINOR + 1))
    PATCH=0
    ;;
  patch)
    PATCH=$((PATCH + 1))
    ;;
  *)
    echo "Invalid bump type: $BUMP_TYPE"
    exit 1
    ;;
esac

NEW_VERSION="$MAJOR.$MINOR.$PATCH"
echo "New version: $NEW_VERSION"

echo "$NEW_VERSION" > VERSION

TODAY=$(date +%Y-%m-%d)
CHANGELOG_ENTRY="## v$NEW_VERSION ($TODAY)\n\n"

process_devrev_changelog() {
  local changelog_dir=$1
  local section_title=$2

  if [ -d "$changelog_dir" ]; then
    local recent_files=$(ls -1 "$changelog_dir"/*.md 2>/dev/null | sort -r | head -5)

    if [ -n "$recent_files" ]; then
      CHANGELOG_ENTRY+="### $section_title\n\n"

      for file in $recent_files; do
        local date=$(basename "$file" .md)
        local content=$(cat "$file")

        if ! grep -q "included from $date" CHANGELOG.md 2>/dev/null; then
          CHANGELOG_ENTRY+="#### Changes from $date\n\n"
          CHANGELOG_ENTRY+="$content\n\n"
          CHANGELOG_ENTRY+="<!-- included from $date -->\n\n"
        fi
      done
    fi
  fi
}

process_devrev_changelog "changelogs/public" "Public API Changes"
process_devrev_changelog "changelogs/beta" "Beta API Changes"

if [ "$CHANGELOG_ENTRY" == "## v$NEW_VERSION ($TODAY)\n\n" ]; then
  CHANGELOG_ENTRY+="### Changes\n\n"
  CHANGELOG_ENTRY+="- Updated OpenAPI specifications\n"
  CHANGELOG_ENTRY+="- Regenerated SDK clients\n\n"
fi

if [ -f CHANGELOG.md ]; then
  echo -e "$CHANGELOG_ENTRY$(cat CHANGELOG.md)" > CHANGELOG.md
else
  echo -e "# Changelog\n\n$CHANGELOG_ENTRY" > CHANGELOG.md
fi

echo "Version bumped to $NEW_VERSION"