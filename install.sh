#!/usr/bin/env sh

if command -v python3 &>/dev/null; then
    python3 ./install.py $1
else
    echo Python 3 is not installed
    exit 1
fi

