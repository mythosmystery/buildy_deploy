#! /bin/bash

git pull
go build -o server ./...

# Restart the server
sudo systemctl restart buildy_deploy.service
