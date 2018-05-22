#!/usr/bin/env sh

set -e

platform='unknown'
unamestr=`uname`

if [[ "$unamestr" == 'Linux' ]]; then
   platform='linux'
elif [[ "$unamestr" == 'FreeBSD' ]]; then
   platform='freebsd'
else
    echo "Plateform not supported"
    exit 1
fi

wget -O /tmp/lifesaver "https://raw.githubusercontent.com/victorboissiere/lifesaver/go/releases/$plateform"
rm /tmp/lifesaver
