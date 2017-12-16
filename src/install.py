#!/usr/bin/env python3

from sys import argv
import json
import subprocess

with open('config.json') as data_file:
    INSTALL_CONFIG = json.load(data_file)


def print_usage():
    print("Usage: ./install.sh [INSTALL_MODE]\n")
    print("Please choose a valid install mode.\nAvailable modes:")
    for mode, install_options in INSTALL_CONFIG.items():
        print("\t- {0}: {1}".format(mode, install_options["description"]))

    exit(1)


def get_shell_command(package_manager, shell_command):
    command_substitutions = {
        "[PKG]": package_manager,
    }

    for str_substitution, substitutionValue in command_substitutions.items():
        shell_command.replace(str_substitution, substitutionValue)

    return shell_command


def install(mode, install_settings):
    print("{0} installation configuration...\n".format(mode.upper()))
    print(install_settings["description"])

    for program in install_settings["programs"]:
        subprocess.check_call("{0} install {1}".format(packageManager, program), shell=True)

    for step in install_mode["steps"]:
        print(step["description"])

        for command in step["commands"]:
            subprocess.check_call(get_shell_command(command), shell=True)


if __name__ == "__main__":
    if len(argv) != 3 or argv[2] not in INSTALL_CONFIG:
        print_usage()

    (_, packageManager, install_mode) = argv
    settings = INSTALL_CONFIG[install_mode]

    if "dependencies" in settings:
        for dependency in INSTALL_CONFIG[install_mode]["dependencies"]:
            install("[DEP] {0}".format(dependency), INSTALL_CONFIG[dependency])

    install(install_mode, INSTALL_CONFIG[install_mode])


