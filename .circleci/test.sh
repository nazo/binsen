#!/bin/bash

set -eux

PULL_REQUEST_NUM=`echo $CIRCLE_PULL_REQUEST | sed -e 's/^[^0-9]*//'`

cd /server/src/github.com/nazo/binsen/server
go test -coverprofile=coverage.txt -covermode=atomic -v ./...
bash <(curl -s https://codecov.io/bash) -cF go -B ${CIRCLE_BRANCH} -C ${CIRCLE_SHA1} -b ${CIRCLE_BUILD_NUM} -P ${PULL_REQUEST_NUM}

cd /client
npm run lint
npm test
bash <(curl -s https://codecov.io/bash) -cF javascript -B ${CIRCLE_BRANCH} -C ${CIRCLE_SHA1} -b ${CIRCLE_BUILD_NUM} -P ${PULL_REQUEST_NUM}
