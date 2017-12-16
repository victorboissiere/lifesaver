# LifeSaver

LifeSafer is a tool to get started on any shell environment (vps, docker, etc.)
by installing only the tools that you need. 

**Work in progress...**

# Getting started

With curl
```bash
curl https://raw.githubusercontent.com/victorboissiere/lifesaver/master/install.sh -fsSL | sh -s minimal
```

```bash
wget -O - https://raw.githubusercontent.com/victorboissiere/lifesaver/master/install.sh | sh -s minimal
```

The `basic` keyword is an installation mode. Check all installation modes
in `config.json`.

# Configuration

To customize available install mode and the pograms you want to install,
simply modify the `config.json` file.

## Config file example

```json
{
 "minimal": {
    "description": "Install basic shell configuration",
    "programs": ["vim", "curl"],
    "steps": [
      {
        "description": "Installing vim configuration",
        "commands": [
          "curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim",
          "cp ./softwares/vim/.vimrc ~/.vimrc"
        ]
      }
    ]
  }
}
```

`programs` and `steps` are optionals

## Dependencies

If you want to have an installation mode that inherit from another, you can
use `dependencies`.

Example:
```json
{
 "full": {
    "description": "Install full shell configuration",
    "dependencies": ["minimal", "shell"]
  }
}
```

This will be the same as running `./install minimal && ./install shell`.
