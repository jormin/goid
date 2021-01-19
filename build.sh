#!/bin/sh
# 编译linux平台
cd cmd/tcp && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build id.go
# 打包镜像
cd ../../ && docker build -t ccr.ccs.tencentyun.com/jormin/goid:latest .
# 上传镜像
docker push ccr.ccs.tencentyun.com/jormin/goid:latest
# 删除生成的可执行文件
rm cmd/tcp/id