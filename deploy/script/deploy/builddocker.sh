#!/bin/bash
# 放置在执行文件目录上一级

IMAGE_NAME=${PWD##*/}
echo "IMAGE_NAME:$IMAGE_NAME:latest"
echo "start bulid..."
docker build -t $IMAGE_NAME:latest .
echo "image success!"

echo "## END ##"
