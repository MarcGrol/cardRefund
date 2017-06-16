#!/bin/bash

if [ "$(git diff --name-status)" ]
then
    echo "Local modifications:"
    git diff --minimal
    exit -1
fi
