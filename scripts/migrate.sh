#!/bin/bash

echo "Cleanup"
rm -rf ./build/migrate

echo "Building migrate tool"
mkdir -p ./build/migrate
go build -o ./build/migrate/migrate ./migrate.go

sleep 1
echo "Signing binary"
codesign --force --sign "mad_server_v1" ./build/migrate/migrate

sleep 1
echo "Running migrate tool"
./build/migrate/migrate