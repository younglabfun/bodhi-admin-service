#!/usr/bin/env bash

echo "build rpc..."
goctl rpc protoc admin.proto --go_out=. --go-grpc_out=. --zrpc_out=../ --style=goZero -m
sed -i "" 's/,omitempty//g' ./admin/*.pb.go
