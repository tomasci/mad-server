#!/bin/bash

echo "Cleanup"
rm -rf ./build/server

echo "Building server"
mkdir -p ./build/server
go build -o ./build/server/server ./main.go

sleep 1
echo "Signing binary"
codesign --force --sign "mad_server_v1" ./build/server/server