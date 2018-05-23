# LifeSaver

[![CircleCI](https://circleci.com/gh/victorboissiere/lifesaver.svg?style=svg)](https://circleci.com/gh/victorboissiere/lifesaver)

LifeSafer is a tool to get started on any shell environment (vps, docker, etc.)
by installing only the tools that you need.

# Getting started

With curl
```bash
curl https://ls.gitcommit.fr -fsSL | bash -s minimal $USER
```

```bash
wget -O - https://ls.gitcommit.fr | bash -s minimal $USER
```

If you do not trust the `https://ls.gitcommit.fr`, it is just a simple redirection to `https://raw.githubusercontent.com/victorboissiere/lifesaver/master/install.sh`.

The `minimal` keyword is an installation mode. Check all installation modes
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

You can safely assume that `pip3` and `python3` is available on the host.

## Dependencies

If you want to have an installation mode that inherit from another, you can
use `dependencies`.

Example:

```yaml
full:
  description: Install full shell configuration
  dependencies:
    - minimal
    - shell
```

This will be the same as running `./src/install.py minimal && ./src/install.py shell`.

# Special variables substitutions

- CP: `cp -a`
- PKG: detected package manager
- CHOWN: `chown -R $USER:$USER`
- USER: `$USER`

Where `$USER` can be overriden as the parameter given in second argument (bash -s minimal my_custom_user)
