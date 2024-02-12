#!/usr/bin/env bash

clear

if [[ -e ./main ]]; then 
    rm -rf ./main
fi

go build -o ./main

./main

if [[ -e ./main ]]; then 
    rm -rf ./main
fi