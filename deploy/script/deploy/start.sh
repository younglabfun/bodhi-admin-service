#!/bin/sh
# 与docker-compose配置文件同级
echo "stop service ..."
docker-compose down
echo "start service ..."
docker-compose up -d

echo "start service finished !!!"
