#!/usr/bin/env bash

pacman -Syyy

if [ ! "$(command -v git)" ]; then
    pacman -S git
fi

if [ "$(command -v git)" ]; then
    echo "Git is installed"
fi

git clone https://github.com/MiraiMindz/ArchInstall.git


sh ./ArchInstall/install.sh

