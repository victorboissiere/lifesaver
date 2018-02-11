# LifeSaver

[![Build Status](https://travis-ci.org/victorboissiere/lifesaver.svg?branch=master)](https://travis-ci.org/victorboissiere/lifesaver)

LifeSafer is a tool to get started on any shell environment (vps, docker, etc.)
by installing only the tools that you need.

# Getting started

With curl
```bash
curl https://raw.githubusercontent.com/victorboissiere/lifesaver/master/install.sh -fsSL | bash -s minimal $USER
```

```bash
wget -O - https://raw.githubusercontent.com/victorboissiere/lifesaver/master/install.sh | bash -s minimal $USER
```

The `basic` keyword is an installation mode. Check all installation modes
in `config.yaml`.

# Configuration

To customize available install mode and the pograms you want to install,
simply modify the `config.yaml` file.

## Config file example

```yaml
minimal:
  description: Install basic shell configuration
  programs:
    - vim
    - curl
    - zsh
  steps:
    - description: Installing vim configuration
      commands:
        - [CP] ./softwares/vim/.vimrc ~/.vimrc
        - curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
        - ["vim", "+silent", "+PlugInstall", "+qall"]
        - [CHOWN] ~/.vim
```

`programs` and `steps` are optionals

## Dependencies

If you want to have an installation mode that inherit from another, you can
use `dependencies`.

Example:

```yaml
full:
  description: Install full shell configuration
  dependencies:
    - miniaml
    - shell
```

This will be the same as running `./src/install.py minimal && ./src/install.py shell`.

# Special varibles substitutions

- CP: `cp -a`
- PKG: detected package manager
- CHOWN: `chown -R $USER:$USER`
- USER: `$USER`

Where $USER can be overriden as the parameter given in second argument (bash -s minimal my_custom_user)
