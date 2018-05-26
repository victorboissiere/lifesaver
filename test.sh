#!/usr/bin/env bash

set -e

if [ $# -eq 1 ] && [ "$1" = "build" ]; then
  docker build -t lifesaver .
fi

docker run --rm --name lifesaver lifesaver /tmp/lifesaver
