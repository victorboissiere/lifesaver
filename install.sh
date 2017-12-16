#!/usr/bin/env sh

set -e

if [ -f /etc/alpine-release ]; then
    PACKAGE_MANAGER='apk'
if [ -f /etc/alpine-release ]; then
    PACKAGE_MANAGER='apk'
else
    echo Could not detect OS
    exit 1
fi

for soft in python3 git
do
    PACKAGE_MANAGER install ${soft}
done

if command -v python3 &>/dev/null; then
    python3 ./install.py ${PACKAGE_MANAGER} $1
fi

