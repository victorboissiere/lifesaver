#!/usr/bin/env bash

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


echo "Please provide sudo pass for software install (simple sudo echo trigger):"
sudo echo -e "Thanks! All set!\n"
echo "Downloading latest binary version $(wget -q -O - https://api.github.com/repos/victorboissiere/lifesaver/releases/latest  | grep tag_name | cut -d '"' -f 4)"
wget -q -O /tmp/lifesaver "$(wget -q -O - https://api.github.com/repos/victorboissiere/lifesaver/releases/latest | grep browser_download_url | cut -d '"' -f 4)"
echo -e "Done!\n"

chmod +x /tmp/lifesaver
/tmp/lifesaver $@
rm /tmp/lifesaver
