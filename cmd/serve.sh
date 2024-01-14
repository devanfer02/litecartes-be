#!/bin/bash

if [ -z "$1" ]; then 
    go run app/main.go 
else 
    go run app/main.go generate
fi