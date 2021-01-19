FROM alpine:latest

COPY ./cmd/tcp/id /usr/bin

EXPOSE 8888

CMD ["/usr/bin/id"]