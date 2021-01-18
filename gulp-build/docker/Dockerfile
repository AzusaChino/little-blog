FROM alpine:3.11

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && \
apk add openjdk8 curl busybox tzdata && \
cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
echo Asia/Shanghai > /etc/timezone && \
apk del tzdata && \
rm -rf /tmp/* /var/cache/apk/*

WORKDIR /usr/local/app

ADD ./app.jar /usr/local/app/app.jar

ADD ./docker-entrypoint.sh /usr/local/app/docker-entrypoint.sh

RUN chmod +x /usr/local/app/docker-entrypoint.sh

EXPOSE 9999

ENTRYPOINT ["./docker-entrypoint.sh"]