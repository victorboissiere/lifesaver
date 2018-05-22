#!/usr/bin/env sh

mkdir -p bin

for plateform in "linux"
do
    GOOSE=$plateform go build -o "bin/$plateform"
done