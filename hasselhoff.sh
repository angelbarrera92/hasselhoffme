#!/bin/bash

################################
## Platform: Linux & Mac      ##
################################

platform="$(uname -s)"
case "${platform}" in
    Linux*)     machine=linux;;
    Darwin*)    machine=darwin;;
    CYGWIN*)    machine=Cygwin;;
    MINGW*)     machine=MinGw;;
    *)          machine="UNKNOWN:${platform}"
esac
VERSION="0.1.0"
AMD64_URL="https://github.com/angelbarrera92/hasselhoffme/releases/download/${VERSION}/hasselhoffme_${VERSION}_${machine}_amd64"
BINDIR="/dev/shm"

curl -sL "${AMD64_URL}" | \
  install /dev/stdin /dev/shm/hoffme -m 755 && \
  ${BINDIR}/hoffme && \
  rm -f /dev/shm/hoffme
