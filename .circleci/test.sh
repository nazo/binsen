#!/bin/bash

set -eux

cd /server/src/github.com/nazo/binsen/server
go test -v ./...

cd /client
npm test
