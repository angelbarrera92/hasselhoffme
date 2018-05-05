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

VERSION="0.0.2"
AMD64_URL="https://github.com/angelbarrera92/hasselhoffme/releases/download/${VERSION}/hasselhoffme_${VERSION}_${machine}_amd64.tar.gz"

## Local directory for binary
BINDIR="/tmp/david"

## Creating directory if not present ##
mkdir -m a=rwx -p $BINDIR

## Download the binary to local path ##
cd $BINDIR && curl -s -L $AMD64_URL | tar xz

## Setting desktop background image from using the binary ##
$BINDIR/hasselhoffme