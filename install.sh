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

export SUDO_UID
export SUDO_GID
wget -O /tmp/lifesaver "https://raw.githubusercontent.com/victorboissiere/lifesaver/go/bin/$platform"
chmod +x /tmp/lifesaver
sudo /tmp/lifesaver $@
rm /tmp/lifesaver
