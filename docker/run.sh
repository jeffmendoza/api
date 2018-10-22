#!/bin/ash

set -e

cd istio.io/api/

/usr/bin/protolock status

cd ../../

/usr/bin/protoc -I/protobuf "$@"
