FROM alpine:latest
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

#定义时区参数
ENV TZ=Asia/Shanghai
#安装时区数据包
RUN apk add tzdata
#设置时区
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo '$TZ' > /etc/timezone

# 在容器根目录 创建一个 apps 目录
WORKDIR /opt/admin


# 拷贝执行文件
# 拷贝文件
COPY ./admin-api /opt/admin/

# 拷贝配置文件到容器中
COPY etc/adminapi-api.yaml /opt/admin/etc/

# 设置编码
ENV LANG C.UTF-8

# 暴露端口
EXPOSE 8888

# 运行golang程序的命令
ENTRYPOINT ["/opt/admin/admin-api"]
