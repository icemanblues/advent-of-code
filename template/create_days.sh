#!/bin/bash

# Sets up an advent-of-code directory structure
# Must be run in the root directory

if [ -z "$1" ]; then
    echo "first param is advent year"
    exit 1
fi

if [ -z "$2" ]; then
    echo "second param is the programming language"
    exit 2
fi

mkdir -p $1

for i in {01..25}
do
    mkdir -p $1/day${i}
    cp template/dayXX.$2 $1/day${i}/day${i}.$2
done
