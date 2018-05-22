#!/usr/bin/env sh

mkdir -p bin

for plateform in "linux"
do
    GOOS=linux go build -o "bin/$plateform"
done