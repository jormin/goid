pipeline{
    agent any
    stages {
        stage('Compile'){
            steps{
                sh """
                    # 更改库获取方式
                    git config --global url."ssh://git@gitlab.wcxst.com:2224".insteadOf "https://gitlab.wcxst.com"
                    git config --global url."ssh://git@gitlab.wcxst.com:2224".insteadOf "http://gitlab.wcxst.com"
                    # 设置私库。如果要永久有效请配置到环境变量
                    export GOPRIVATE=gitlab.wcxst.com
                    # 配置国内地址
                    export GOPROXY=https://goproxy.io
                    # 下载依赖
                    /usr/local/go/bin/go mod download
                    # 处理配置文件
                    rm -f config.yaml
                    cp config-prod.yaml config.yaml
                    # 编译
                    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 /usr/local/go/bin/go build cmd/tcp/id.go
                """
            }
        }
        stage('Build'){
            steps{
                sh """
                    docker build -t ccr.ccs.tencentyun.com/jormin/goid:latest .
                """
            }
        }
        stage('Push'){
            steps{
                sh """
                    docker push ccr.ccs.tencentyun.com/jormin/goid:latest
                """
            }
        }
        stage('Deploy'){
            steps{
                sh """
                    ssh -p 22 root@82.157.101.82 "kubectl rollout restart deploy goid -n=jormin"
                """
            }
        }
    }
}
