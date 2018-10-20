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
VERSION="0.2.0"
AMD64_URL="https://github.com/angelbarrera92/hasselhoffme/releases/download/${VERSION}/hasselhoffme_${VERSION}_${machine}_amd64"
BINDIR="/tmp/david"
mkdir -m a=rwx -p $BINDIR
cd $BINDIR && curl -s -L $AMD64_URL -O
chmod u=rwx $BINDIR/hasselhoffme_${VERSION}_${machine}_amd64
if [ -z "$DISPLAY" ] && [ "$machine" != "darwin" ];then
    sudo $BINDIR/hasselhoffme_${VERSION}_${machine}_amd64 setmotd
else
    $BINDIR/hasselhoffme_${VERSION}_${machine}_amd64
fi
