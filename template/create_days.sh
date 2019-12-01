#!/bin/bash

#
# Sets up an advent-of-code directory structure
#

# TODO: should prompt or cmd args on which programming language
# typescript assumed
# TODO: should prompt or cmd args on which event (2019) and where does it live
# 2019 assumed, one directory up

for i in {02..25}
do
    mkdir -p day${i}
    cp dayXX.ts day${i}/day${i}.ts
    mv day${i} ../2019
done
