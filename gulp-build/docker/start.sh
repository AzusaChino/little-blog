#!/usr/bin/env sh
docker run -d -e ZK_HOST = 127.0.0.1:2181 -p 8888:9999 --name=xxx xxx:1.0.0