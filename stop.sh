#!/bin/sh

ps ax|grep ability | grep -v grep | awk -F' ' '{print $1}' | xargs kill 

