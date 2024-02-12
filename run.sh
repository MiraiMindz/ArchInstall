#!/usr/bin/env bash

setxkbmap -model abnt2 -layout br -variant abnt2
clear
goimports -w ./*/*.go
go vet ./...
# go run $1
go run main.go
