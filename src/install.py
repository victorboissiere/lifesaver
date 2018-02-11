#!/usr/bin/env python3

from sys import argv
import yaml
import getpass
import subprocess

with open('config.yml') as data_file:
    INSTALL_CONFIG = yaml.load(data_file)


def print_usage():
    print("Usage: ./install.sh [INSTALL_MODE]\n")
    print("Please choose a valid install mode.\nAvailable modes:")
    for mode, install_options in INSTALL_CONFIG.items():
        print("\t- {0}: {1}".format(mode, install_options["description"]))

    exit(1)


def get_shell_command(user, package_manager, shell_command):
    command_substitutions = {
        "PKG": package_manager,
        "CP": "cp -a",
        "CHOWN": "chown -R {0}:{0}".format(user),
        "USER": user,
    }

    for str_substitution, substitutionValue in command_substitutions.items():
        shell_command = shell_command.replace(str_substitution, substitutionValue)

    return shell_command


def install(user, package_manager, mode, install_settings):
    print("{0} installation configuration...\n".format(mode.upper()))
    print(install_settings["description"])

    if "programs" in install_settings and len(install_settings["programs"]) > 0:
        packages = " ".join(install_settings["programs"])
        print("[PROGRAMS] installing {0}...".format(packages))
        subprocess.check_call("{0} {1}".format(package_manager, packages), shell=True)

    if "steps" in install_settings:
        for step in install_settings["steps"]:
            print("[STEP] {0}".format(step["description"]))

            for command in step["commands"]:
                if isinstance(command, list):
                    subprocess.check_call(command)
                else:
                    subprocess.check_call(get_shell_command(user, package_manager, command), shell=True)


if __name__ == "__main__":
    if len(argv) < 3 or argv[2] not in INSTALL_CONFIG:
        print_usage()

    userInstall = getpass.getuser()
    if len(argv) == 3:
        (_, packageManager, install_mode) = argv
    else:
        (_, packageManager, install_mode, userInstall) = argv

    settings = INSTALL_CONFIG[install_mode]

    if "dependencies" in settings:
        for dependency in INSTALL_CONFIG[install_mode]["dependencies"]:
            install(userInstall, packageManager, "[DEP] {0}".format(dependency), INSTALL_CONFIG[dependency])

    install(userInstall, packageManager, install_mode, INSTALL_CONFIG[install_mode])
