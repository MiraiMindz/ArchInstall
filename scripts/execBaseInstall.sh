#!/usr/bin/env bash

SCRIPTS_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
chmod +x $SCRIPTS_DIR/BaseInstall
$SCRIPTS_DIR/BaseInstall
