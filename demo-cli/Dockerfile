FROM alpine:latest

#更新安装源
RUN sed -i 's/dl-cdn.alpineliunx.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk add --no-cache bash

RUN mkdir -p /app
WORKDIR /app

ADD demo-cli /app/demo-cli

CMD ["./demo-cli"]