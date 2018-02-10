#!/usr/bin/env bash

set -e

cp ~/.vimrc ./softwares/vim/
cp ~/.zshrc ./softwares/zsh

if [  -n "$(git status --porcelain)" ]; then
  echo "Done!"
  echo "Please commit your changes"
else
  echo "No changes detected"
fi
