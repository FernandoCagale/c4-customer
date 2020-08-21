#!/bin/sh

set -e

[ -z "$DEBUG" ] || set -x

echo "\n===> Generate image...\n"

docker build --no-cache -f Dockerfile.http -t c4-customer .

echo "\n===> Docker tag...\n"

docker tag c4-customer fernandocagale/c4-customer:http

echo "\n===> Docker push...\n"

docker push fernandocagale/c4-customer:http