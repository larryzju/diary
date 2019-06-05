#/usr/bin/env bash

set -e

REPOPATH=$(git rev-parse --show-toplevel)
REPONAME=$(basename ${REPOPATH})
TEMP=$(mktemp -d)
pushd ${TEMP}
git clone ${REPOPATH}
