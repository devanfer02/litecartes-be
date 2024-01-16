#!/bin/bash

python="python"
config_dir="config"
sdk_file="$config_dir/litecartes-firebase-sdk.json"

echo "creating config directory"
mkdir -p "$config_dir"

echo "creating python directory"
mkdir -p "$python"
mkdir -p "$python/config"

touch "$python/config/litecartes-glcoud.json"
echo '{"desc": "put gcloud private key in here to use vertex api"}' 

echo "creating firebase-sdk.json file in config directory"
touch "$sdk_file"
echo '{"desc": "put firebase private key in here"}' > "$sdk_file"

echo "creating .env file"
cp ".env.example" ".env"

echo "downloading application needed dependancies"
go mod download

echo "initialization complete."
echo "proceed to configure .env, firebase-sdk.json, litecartes-glcoud.json and run the server"
