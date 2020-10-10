#!/bin/sh

git pull
go build
./restart.sh
