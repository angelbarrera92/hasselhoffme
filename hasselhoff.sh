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
arch="$(uname -m)"
case "${arch}" in
    armv6l)    arch=armv6;;
    armv7l)    arch=armv7;;
    *)         arch="amd64"
esac
VERSION="0.4.0"
BIN_URL="https://github.com/angelbarrera92/hasselhoffme/releases/download/${VERSION}/hasselhoffme_${VERSION}_${machine}_${arch}"
BINDIR=$(mktemp -d)
cd $BINDIR && curl -s -L $BIN_URL -O
chmod u=rwx $BINDIR/hasselhoffme_${VERSION}_${machine}_${arch}
if [ -z "$DISPLAY" ] && [ "$machine" != "darwin" ];then
    sudo $BINDIR/hasselhoffme_${VERSION}_${machine}_${arch} setmotd
else
    $BINDIR/hasselhoffme_${VERSION}_${machine}_${arch}
fi
rm -rf $BINDIR
