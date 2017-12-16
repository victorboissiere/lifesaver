#!/usr/bin/env sh

set -e

if [ -f /etc/alpine-release ]; then
    PACKAGE_MANAGER='apk'
elif [ -f /etc/debian_version ]; then
    apt-get update
    PACKAGE_MANAGER='apt-get -y'
else
    echo Could not detect OS
    exit 1
fi

for soft in python3 git
do
    if command -v ${soft} &>/dev/null; then
        echo "Installing ${soft} required package"
        ${PACKAGE_MANAGER} install ${soft}
    fi
done

rm -rf /tmp/lifesaver
git clone https://github.com/victorboissiere/lifesaver /tmp
cd /tmp/lifesaver && python3 ./src/install.py "${PACKAGE_MANAGER}" $1

