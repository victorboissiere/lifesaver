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


echo "Download latest binary..."
wget -q -O /tmp/lifesaver "$(wget -q -O - https://api.github.com/repos/victorboissiere/lifesaver/releases/latest | grep browser_download_url | cut -d '"' -f 4)"
echo -n "Done!\n"
chmod +x /tmp/lifesaver
sudo SUDO_UID=${SUDO_UID} SUDO_GID=${SUDO_GID} /tmp/lifesaver $@
rm /tmp/lifesaver
