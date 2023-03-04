#!/usr/bin/env bash


setxkbmap -model abnt2 -layout br -variant abnt2
clear
go fmt $1
go run $1
