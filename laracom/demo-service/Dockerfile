# 使用alpine
# 第一行 FROM 指令表示基于 Linux Alpine 的最新版作为基础镜像，
# Alpine 是一个面向安全的轻型 Linux 发行版，
# 不同于普通的 Linux 发行版，Alpine 采用了 musl libc 和 busybox 以减小系统的体积和运行时资源消耗，
# Alpine Docker 镜像也继承了 Alpine Linux 发行版的这些优势，相比于其他 Docker 镜像，
# 它的体积非常小，仅仅只有 5 MB 左右（而 Ubuntu 系列镜像通常会接近 200 MB）。
FROM alpine:latest

# 第二行指令用于更新 Alpine 软件镜像源（推荐阿里、清华或中科大的镜像源），
# 在国内各种 Docker 镜像构建，这都是必不可少的步骤。
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 接下来我们更新 Alpine 的包管理器 apk，并安装 bash 和 supervisor，前者用于执行 shell 命令，后者用于管理多个服务的自启动。
RUN apk update && apk add --no-cache bash supervisor

# 在 Docker 镜像中创建一个 /app 目录
RUN mkdir /app

# WORKDIR指令用于指定容器的一个目录，容器启动时执行的命令会在该目录下执行，相当于设置了容器的工作目录。
WORKDIR /app

# 然后把 Dockerfile 所在目录下的 demo-service 二进制文件（位于宿主机，稍后会通过运行 go build 命令生成）拷贝到镜像的 /app 目录下
ADD demo-service.exe /app/demo-service

# 把当前目录下的 supervisord.conf 配置文件拷贝到镜像的 /etc/ 目录下。
ADD supervisord.conf /etc/supervisord.conf

# 最后一行 CMD 指令表示容器启动时默认执行的命令，这里我们启动 Supervisor，并通过它来启动和管理 gRPC 和 HTTP 服务。
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]