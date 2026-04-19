# 使用明确的版本号，比 latest 更稳定
FROM alpine:3.19

# 1. 设置环境变量 (合并设置)
ENV TZ=Asia/Shanghai \
    LANG=C.UTF-8 \
    APP_HOME=/opt/admin

# 2. 合并系统操作：换源、安装时区插件、创建用户、清理缓存
# - 使用 --no-cache 不保存 apk 索引，进一步缩小体积
# - 创建非 root 用户运行程序，提升安全性
RUN set -eux && \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk add --no-cache tzdata && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo "$TZ" > /etc/timezone && \
    addgroup -S admin && adduser -S admin -G admin

# 3. 设置工作目录
WORKDIR $APP_HOME

# 4. 拷贝文件并修改所有者 (一步到位)
# 建议先拷贝配置文件，再拷贝频繁变动的二进制文件
COPY --chown=admin:admin ./rpc/etc/admin.yaml ./etc/
COPY --chown=admin:admin ./rpc/admin-rpc ./

# 5. 切换到非 root 用户
USER admin

# 6. 暴露端口
EXPOSE 9020

# 7. 运行命令
# 使用绝对路径更稳妥
ENTRYPOINT ["./admin-rpc"]
# 如果需要指定配置文件路径，可以配合 CMD 使用：
 CMD ["-f", "etc/admin.yaml"]
