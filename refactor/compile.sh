#!/usr/bin/env bash
export GOOS=linux
export GOARCH=amd64

GOOS=linux
GOARCH=amd64

if [ -f "./scripts/BaseInstall" ]; then
    rm -v "./scripts/BaseInstall"
fi

if [ -f "./scripts/PreInstall" ]; then
    rm -v "./scripts/PreInstall"
fi

if [ -f "./scripts/Startup" ]; then
    rm -v "./scripts/Startup"
fi


cd ./sections/BaseInstall || exit
go build -compiler=gc .
cd ../../

cd ./sections/PreInstall || exit
go build -compiler=gc .
cd ../../

cd ./sections/Startup || exit
go build -compiler=gc .
cd ../../

mv -v ./sections/Startup/Startup ./scripts/Startup
mv -v ./sections/PreInstall/PreInstall ./scripts/PreInstall
mv -v ./sections/BaseInstall/BaseInstall ./scripts/BaseInstall

