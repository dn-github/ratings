#!/bin/sh

set -e
export GO111MODULE=on
GOOS=linux go build -o ratings
docker build -t devnryn/ratings:v1 .
docker push devnryn/ratings:v1
rm ratings
