#!/usr/bin/env bash

# PATHS
# Get current working directory
current_dir=`pwd`

# Get the absolute path of where script is running from
script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd)"
script_path="$script_dir/run.sh"

# RETURN VARIABLE
ret=""

# ARGUMENTS
args="${@:2}" # All arguments except the first

# DESCRIPTION:
#   Where execution starts
main() {
    case $1 in
        install )
            install $2
        ;;
        --help|help|-h )
            help
        ;;
    esac

    exit 0
}

# TODO: Debug install (wasmod) vs release install (wasmo)
# DESCRIPTION:
#   Installs wasmo project
install() {
    echo "A work in progress"
}

help() {
    echo "A work in progress"
}

# Start main
main $@
