#!/usr/bin/env python3

from sys import argv
import json
import subprocess

with open('config.json') as data_file:
    INSTALL_CONFIG = json.load(data_file)

if len(argv) != 3:
    print("Please choose a valid install mode. Available modes:\n- {0}".format("\n- ".join(INSTALL_CONFIG.keys())))
    exit(1)

(_, packageManager, installMode) = argv

COMMAND_SUBSTITUTIONS = {
    "[PKG]": packageManager,
}


def get_shell_command(cmd):
    shell_command = cmd
    for str_substitution, substitutionValue in COMMAND_SUBSTITUTIONS.items():
        shell_command.replace(str_substitution, substitutionValue)

    return shell_command


install_mode = INSTALL_CONFIG[installMode]
print(install_mode["description"])

for program in install_mode["programs"]:
    subprocess.check_call("{0} install {1}".format(packageManager, program))

for step in install_mode["steps"]:
    print(step["description"])

    for command in step["commands"]:
        subprocess.check_call(get_shell_command(command), shell=True)


