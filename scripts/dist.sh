#!/usr/bin/env bash
set -e

# Get the version from the command line
VERSION=$1
if [ -z "$VERSION" ]; then
    echo "Please specify a version."
    exit 1
fi

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that dir because we expect that
cd "$DIR"

# Zip all the files
rm -rf ./pkg/dist
mkdir -p ./pkg/dist
for FILENAME in $(find ./pkg -mindepth 1 -maxdepth 1 -type f); do
    FILENAME="$(basename "${FILENAME}")"
    cp "./pkg/${FILENAME}" "./pkg/dist/arukas_${VERSION}_${FILENAME}"
done

ghr -t ${GITHUB_TOKEN} -u ${USERNAME} -r ${CIRCLE_PROJECT_REPONAME} --replace ${VERSION} dist/

exit 0
