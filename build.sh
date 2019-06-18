#!/usr/bin/env bash

set -ex

echo $PWD
DEFAULT_REPO=https://github.com/larryzju/diary
REPO=$1
REPO=${REPO:-${DEFAULT_REPO}}
MASTER_WORKTREE=_master

if [ ! -d ${MASTER_WORKTREE} ]; then
	git worktree add -f ${MASTER_WORKTREE} master
fi

MASTER_COMMIT=$(cd $PWD/${MASTER_WORKTREE}; git pull ${REPO}> /dev/null && git rev-parse HEAD)
LAST_BUILD_COMMIT=$(cat BUILD.sha1)
if [ ${MASTER_COMMIT} == ${LAST_BUILD_COMMIT} ]; then 
	echo "no more commit"
	exit 0
fi

swgen -input ${MASTER_WORKTREE} -output $PWD -verbose -root /diary -template ${MASTER_WORKTREE}/_template/page.html
echo ${MASTER_COMMIT} > BUILD.sha1
git add --all
git commit -m "rebuild: ${LAST_BUILD_COMMIT} -> ${MASTER_COMMIT}"
