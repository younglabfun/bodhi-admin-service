#!/usr/bin/env bash

# 使用方法：
#获取gorm-zero项目
#  git clone https://github.com/SpectatorNan/gorm-zero.git
#使用gorm-zero项目修改template
#  修改template，在项目目录下
#  goctl template init --home ./template
#  使用gorm-zero项目下1.4.2的model替换template目录下model

# ./genModel.sh user
# ./genModel.sh category
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package


#生成的表名
tables=$1
#表生成的genmodel目录
modeldir=./genModel

# 数据库配置
host=127.0.0.1
port=33060
dbname=bodhi-admin
username=root
passwd=123456


echo "开始创建库：$dbname 的表：$2"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=false --style=goZero --home ../../template
