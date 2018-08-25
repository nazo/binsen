#!/bin/bash

set -eux

cd ./server
go test -v ./...
cd ..

cd ./client
npm test
cd ..
