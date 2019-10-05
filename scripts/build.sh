#!/bin/bash

echo 'compiling executable'
go build -o "../bin/taskforce" -i "../cmd/taskforce/main.go"
