# bodhi admin
Copyright 2024 younglabfun (https://younglab.fun/)

Distributed under MIT license.

Bodhi Admin based on go-zero,vue-admin-template



## bodhi admin

bodhi admin 是一个基础的后台管理系统。bodhi admin 本是为菩提辞典( https://bodhidict.com/ )开发的管理后台，在开发过程中改为独立项目开源。该项目包括后端服务、web前端两部分，包含管理后台的基本功能，用户登录、注册，菜单管理以及基于RBAC模型的角色授权控制。

- bodhi admin service: https://github.com/younglabfun/bodhi-admin-service

- bodhi admin web: https://github.com/younglabfun/bodhi-admin-web



## bodhi admin service

这是bodhi admin项目的微服务后端，基于 go-zero 进行开发。下面为 bodhi admin service 开发、使用说明，关于 bodhi admin web前端部分，请移步  https://github.com/younglabfun/bodhi-admin-web

 

### 目录结构

```shell
.
├── LICENSE # license
├── README.md # readme
├── app # 应用目录
│   └── admin # admin service
├── common # 公共库
│   ├── errorx # errorX 错误库
│   ├── responsex # responseX 应答库
│   └── utils
├── deploy # 发布相关
│   ├── script # 脚本
│   │   ├── deploy # 发布脚本，包括docker方式运行配置等
│   │   ├── gencode # goctl脚本
│   │   └── mysql # goctl model脚本
│   ├── sql # sql文件
│   └── template # go-zero模版
├── go.mod
└── go.sum

```



### 运行服务

```bash
# 进入app目录
$ cd /workspace/app/rpc # or api
# 整理依赖文件
$ go mod tidy
# 启动 go 程序
$ go run admin.go  # or adminapi.go
```

- rpc 服务默认运行端口 9020
- api 服务默认运行端口 9021



### 开发脚本

- 生成model文件

  ```bash
  $ cd /workspace/deploy/script/mysql
  $ vi genModel.sh
  # 修改数据库配置
  host=127.0.0.1
  port=33060
  dbname=bodhi-admin
  username=root
  passwd=123456
  cache=false
  
  $ ./genModel.sh table_name
  # 再将./genModel下的文件剪切到对应服务的model目录里面，修改package
  ```

- 自动生成代码

  api定义以及rpc定义详见go-zero文档

  ```bash
  # 生成 rpc
  $ cd /workspace/app/rpc/proto
  $ goctl rpc protoc admin.proto --go_out=. --go-grpc_out=. --zrpc_out=../ --style=goZero -m
  
  # 生成 api
  $ cd /workspace/app/api/desc
  $ goctl api go -api admin.api -dir ../  --style=goZero --home ../../../../deploy/template
  ```



### 编译发布

- 编译执行文件

```bash
# 编译 rpc
$ cd /workspace/app/rpc
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o admin-rpc admin.go

# 编译 api
$ cd /workspace/app/api
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o admin-api adminapi.go

# 可用编译脚本发布
/workspace/deploy/script/deploy/deploy.sh

```



- 发布生产

```bash
# 创建环境必须目录
.
├── bodhi-admin # 生产目录
│   ├── etc # 配置目录
│   │   └── install # 安装 sql 目录，将install.sqlfa
│   ├── admin-rpc
│   └── admin-api
```

docker 方式运行拷贝deploy/script/deploy相关配置文件以及脚本



### 安装

初次运行rpc服务时，自动进入 bodhi admin service 安装流程，通过终端交互完成初始配置。

```bash
$ cd /workspace/app/rpc # 进入rpc目录
$ go mod tidy # 处理依赖文件
$ go run admin.go # 生产环境运行 admin-rpc
Start init rpc conf...
1 . Init base conf...
set base conf is ok!
2 . Init etcd conf...
set etcd service？default no (y/n) # 默认不使用ETCD服务
skip set etcd service...
3 . Init admin conf...
4 . Init auth conf...
5 . Init database conf...
Please enter MySQL host (default 127.0.0.1):
Please enter MySQL port (default 3306):
Please enter MySQL user:root
Please enter MySQL password:123456a
6 . Init redis conf...
set Redis conf？default no (y/n)n # 默认不配置 redis
skip set redis conf...
7 . Init log conf...
Start install database...
Check MySQL conf...
...
```

安装完成会自动创建 rpc 服务配置文件、 api 服务配置文件、安装基础数据库、配置超级管理员账号以及密码等信息。



### 关于作者

技术实践者，热爱技术，具有20年+的开发经验，上过班，创过业，作过独立开发者，全栈，常用技术栈PHP、Python、Golang、JavaScript、CSS、VUE等；生活观察员，骑车、读书、做饭、旅行、写作、探索城市角落，爱大海，不会游泳。

#### 欢迎访问作者网站

- 小实验（作者Blog ）：https://younglab.fun/
- 菩提辞典（佛教典籍搜索）：https://bodhidict.com/

#### 资助 Donate

如果您觉得这个项目帮助到您，可以帮作者买杯咖啡表示鼓励。

<center class="half">
<img src="https://github.com/younglabfun/younglabfun/blob/main/image/wxpay.jpg?raw=true" style="width:260px"/><img src="https://github.com/younglabfun/younglabfun/blob/main/image/alipay.jpg?raw=true" style="width:260px"/>
</center>

- 使用开发中有问题可以发Email到: pagopagi(at)gmail



### License

MIT License

Copyright 2024 younglabfun
