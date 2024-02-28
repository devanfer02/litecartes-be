#!/bin/bash

# deploying api as a service in ubuntu
# make sure to execute

pwd_dir=$(pwd)
curr_dir="${pwd_dir: -13}"
dir_should="litecartes"

if [[ ! $curr_dir == *"$dir_should"* ]]; then 
    echo "current directory should be $dir_should"

    echo "exiting..."
    exit 1
fi 

if [ ! -f "main" ]; then 

    echo "entry file main.exe not found"
    echo "building main.exe"

    main_path="app/main.go"

    if [ ! -f "$main_path" ]; then 
        echo "file in $main_path doesnt exist"

        echo "exiting..."

        exit 1
    fi 

    go build "$main_path"

fi 

echo "creating services"

services_content=$(cat << EOF
[Unit]
Description=Litecartes Service
After=network.target

[Service]
Type=simple
WorkingDirectory=$pwd_dir
ExecStart=$pwd_dir/main

Restart=on-failure
RestartSec=10

StandardOutput=syslog
StandardError=syslog

[Install]
WantedBy=default.target
EOF
)

repo_path="systemd/lite.service"
services_name="litecartes"
services_path="/etc/systemd/system/$services_name.service"

echo "$services_content" > "$repo_path"

cp "$repo_path" "$services_path"

sudo systemctl daemon-reload
sudo systemctl start "$services_name"
sudo systemctl enable "$services_name"

echo "api services created and running"
