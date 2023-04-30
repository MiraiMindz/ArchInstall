#!/usr/bin/env bash

set -a
SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
SCRIPTS_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"/scripts
CONFIG_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"/config
set +a


( bash $SCRIPTS_DIR/execStartup.sh )
( bash $SCRIPTS_DIR/execPreInstall.sh )
( bash $SCRIPTS_DIR/execBaseInstall.sh )
( arch-chroot /mnt /usr/bin/runuser -u $USERNAME -- /home/$USERNAME/ArchInstall/scripts/execUserConfig.sh )