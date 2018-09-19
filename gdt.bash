#!/bin/bash

# Gravity Development Toolkit Wrapper for Linux

if [ -z "$1" ]
then
    echo ""
    echo "Gravity Development Toolkit Wrapper for Linux"
    echo ""
    echo "usage: $0 <command>"
    echo "       $0 genversion"
    echo ""
    exit 1
fi

go run ./developing/tools $1
