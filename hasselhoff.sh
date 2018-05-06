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
echo ${machine}

VERSION="0.0.4"
AMD64_URL="https://github.com/angelbarrera92/hasselhoffme/releases/download/${VERSION}/hasselhoffme_${VERSION}_${machine}_amd64.tar.gz"
BINDIR="/tmp/david"
mkdir -m a=rwx -p $BINDIR
cd $BINDIR && curl -s -L $AMD64_URL | tar xz
$BINDIR/hasselhoffme
