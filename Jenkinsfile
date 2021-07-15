pipeline{
    agent any
    stages {
        stage('Compile'){
            steps{
                sh """
                    # 更改库获取方式
                    git config --global url."ssh://git@github.com:2224".insteadOf "https://github.com"
                    git config --global url."ssh://git@github.com:2224".insteadOf "http://github.com"
                    # 设置私库。如果要永久有效请配置到环境变量
                    export GOPRIVATE=github.com
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
                    docker build -t harbor.local.wcxst.com/jormin/goid:latest .
                """
            }
        }
        stage('Push'){
            steps{
                sh """
                    docker login harbor.local.wcxst.com --username=jormin --password=iYA6caLPRiK7mJ6nLbVs8g==
                    docker push harbor.local.wcxst.com/jormin/goid:latest
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
