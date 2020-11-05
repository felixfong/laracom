#demo-service/Dockerfile

#使用golang官方镜像提供Go运行环境，并且命名为buidler以便后续引用
FROM golang:1.13-alpine as builder

#启用Go Modules 并设置GOPROXY
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

#更新安装源
RUN sed -i 's/dl-cdn.alpineliunx.org/mirrors.aliyun.com/g' /etc/apk/repositories

#安装git
RUN apk --no-cache add git

#设置工作目录
WORKDIR /app/laracom-demo-service

#将当前目录所在的代码拷贝到镜像中
COPY . .

#下载依赖
RUN go mod download

#构建二进制文件，添加一些额外参数以便可以在alpine中运行它
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o laracom-demo-service

#第二阶段的镜像构建
FROM alpine:latest

#更新安装源
RUN sed -i 's/dl-cdn.alpineliunx.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk add --no-cache bash supervisor ca-certificates
RUN mkdir /app
WORKDIR /app
#ADD demo-service /app/demo-service
COPY --from=builder /app/laracom-demo-service/laracom-demo-service .
ADD supervisord.conf /etc/supervisord.conf

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]