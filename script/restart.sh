#!/bin/bash

count=`ps aux | grep 'microgin' | grep -v 'restart.sh' | awk '{print $2}'`

echo $count

if [ -n "$count" ]; then
    ps aux | grep 'microgin' | grep -v 'restart.sh' | awk '{print $2}' | xargs kill
fi

nohup /data/microgin/microgin > /dev/null 2>&1 &
