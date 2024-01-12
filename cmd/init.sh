#!/bin/bash

config_dir="config"
sdk_file="$config_dir/litecartes-firebase-sdk.json"

echo "creating config directory"
mkdir -p "$config_dir"

echo "creating firebase-sdk.json file in config directory"
touch "$sdk_file"
echo '{"desc": "put firebase private key in here"}' > "$sdk_file"

echo "creating .env file"
cp ".env.example" ".env"

echo "downloading application needed dependancies"
go mod download

echo "initialization complete. proceed to configure .env, firebase-sdk.json and run the server"