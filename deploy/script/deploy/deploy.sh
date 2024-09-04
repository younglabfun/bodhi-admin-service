#!/bin/bash

SRVNAME="admin-api"
#启动文件
MAIN_SRC="adminapi.go"

build_bin(){
    echo "build ${SRVNAME}"
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${SRVNAME} ${MAIN_SRC}
}

build_bin
echo "build FINISH"
echo ""
echo "upload ${SRVNAME} -> server"
scp ${SRVNAME} root@server:/root
