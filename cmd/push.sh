#!/bin/bash


if [ -z "$1" ]; then 
    echo "provide type commit"
    exit 1
fi 

if [ -z "$2" ]; then 
    echo "provide message argument"
    exit 1
fi

git add . 
git commit -m "$1"
git push origin main