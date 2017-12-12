#!/usr/bin/env python3
from sys import argv

INSTALL_MODES = {
    "basic": [
        {
            "description": "Installing vim",
            "commands": [
                "sudo apt-get install vim"
            ],
        },
    ],
}

if len(argv) == 1:
    print("Please choose a valid install mode. Available modes:\n- {0}".format("\n- ".join(INSTALL_MODES.keys())))
    exit(1)

mode = argv[1]

for installStep in INSTALL_MODES[mode]:
    print(installStep["description"])



