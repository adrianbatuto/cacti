#!/bin/bash
TAG="$1"
echo "Testing TAG: $TAG"
if ! npx semver "$TAG" > /dev/null 2>&1; then
  echo "Error: Tag '$TAG' is not a valid semantic version."
  exit 1
fi
echo "Tag '$TAG' is valid!"
