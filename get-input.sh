#!/bin/bash

if [ -z "$1" ]; then
    echo "first param is the year"
    exit 1
fi

if [ -z "$2" ]; then
    echo "second param is the day"
    exit 2
fi

# cookies.txt needs to contain your session cookie value
curl -b `cat cookies.txt` https://adventofcode.com/$1/day/$2/input > input.txt
