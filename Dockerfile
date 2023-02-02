# 将项目构建成二进制可执行文件
FROM golang:alpine AS builder

ARG LMS_DIR=/go/src/github.com/lms

WORKDIR $LMS_DIR/server
COPY . .

# 设置go的环境变量，设置后展示go的环境变量列表，下载依赖，进行build构建
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o lms_server .

# 接下来创建镜像
FROM alpine:latest

LABEL MAINTAINER="turbochang@126.com"

ARG LMS_DIR=/go/src/github.com/lms

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置编码
ENV LANG C.UTF-8

# 切换工作目录
WORKDIR $LMS_DIR

# 将build的执行文件、yaml文件从builder的目录下复制到该容器目录下
COPY --from=0 $LMS_DIR/server/lms_server ./
COPY --from=0 $LMS_DIR/server/config-docker.yaml ./

# 挂载，docker会在安装目录下的指定目录下面生成一个目录来绑定容器的匿名卷。方便修改config-docker.yaml文件
VOLUME $LMS_DIR

# 声明服务端口
EXPOSE 8888

# 启动时需要运行的命令
ENTRYPOINT ./lms_server -conf config-docker.yaml