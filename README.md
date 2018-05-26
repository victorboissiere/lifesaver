# LifeSaver

[![CircleCI](https://circleci.com/gh/victorboissiere/lifesaver.svg?style=svg)](https://circleci.com/gh/victorboissiere/lifesaver)

LifeSafer is a tool to get started on any shell environment (vps, docker, etc.)
by installing only the tools that you need.

# Demo

[![asciicast](https://asciinema.org/a/xylgNrCWHdO0x9wfERJzAKwKu.png)](https://asciinema.org/a/xylgNrCWHdO0x9wfERJzAKwKu)

# Getting started

With curl
```bash
curl https://ls.gitcommit.fr -fsSL | bash -s minimal
```

```bash
wget -O - https://ls.gitcommit.fr | bash -s minimal
```

If you do not trust the `https://ls.gitcommit.fr`, it is just a simple redirection to `https://raw.githubusercontent.com/victorboissiere/lifesaver/master/install.sh`.

The `minimal` keyword is an installation mode. Check all installation modes
in `config.yaml`.

