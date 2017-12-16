#!/usr/bin/env bash

set -e
echo $1

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
    which ${soft} >/dev/null && continue || ${PACKAGE_MANAGER} install ${soft}
done

USER=${USER}
if ! [ -z "$2" ]; then
    USER=$2
fi

rm -rf /tmp/lifesaver
cd /tmp && git clone https://github.com/victorboissiere/lifesaver
echo "Executing as user ${USER}"
chown -R ${USER}:${USER} /tmp/lifesaver

cd /tmp/lifesaver && python3 ./src/install.py "${PACKAGE_MANAGER}" $1 ${USER}

rm -rf /tmp/lifesaver

