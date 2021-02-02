FROM alpine:latest

COPY ./id /usr/bin
COPY ./config.yaml /

EXPOSE 8888

RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.7/main/ > /etc/apk/repositories \
    && apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

CMD ["/usr/bin/id"]