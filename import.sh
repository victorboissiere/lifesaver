#!/usr/bin/env bash

set -e

cp ~/.vimrc ./softwares/vim/
cp ~/.zshrc ./softwares/zsh
cp -r ~/.oh-my-zsh/custom ./softwares/zsh
cp ~/.config/ranger/rc.conf ./softwares/ranger
if [ -d ~/.config/terminator ]; then
  cp ~/.config/terminator/config ./softwares/terminator
fi

if [  -n "$(git status --porcelain)" ]; then
  echo "Done!"
  echo "Please commit your changes"
else
  echo "No changes detected"
fi
