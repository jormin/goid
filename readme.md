## About

Go 开发的分布式 ID (雪花ID)服务

#### quick start

```shell script
    # 更改库获取方式
    git config --global url."ssh://git@gitlab.wcxst.com:2224".insteadOf "https://gitlab.wcxst.com"
    git config --global url."ssh://git@gitlab.wcxst.com:2224".insteadOf "http://gitlab.wcxst.com"
    # 设置私库。如果要永久有效请配置到环境变量
    export GOPRIVATE=gitlab.wcxst.com
    # 配置国内地址
    export GOPROXY=https://goproxy.io
    # 下载依赖
    go mod download
```