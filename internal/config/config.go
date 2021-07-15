package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/jormin/gohelper"
	"github.com/jormin/golog/log"
)

const (
	IP     = 8888
	IDCall = "ID.NewID"
)

var (
	// 环境
	Env string

	// 日志
	LogCfg log.Config
)

// 初始化配置
func Init() {
	// 读取配置文件
	path, _ := os.Getwd()
	viper.SetConfigType("yaml")
	viper.SetConfigFile(fmt.Sprintf("%s/config.yaml", path))
	gohelper.Must(viper.ReadInConfig())
	// 初始化配置
	initEnv()
	initLogCfg()
}

// 初始化环境
func initEnv() {
	Env = viper.GetString("env")
}

// 初始化日志配置
func initLogCfg() {
	LogCfg = log.Config{
		App:     viper.GetString("log.app"),
		LogPath: viper.GetString("log.path"),
		LogFmt:  viper.GetString("log.format"),
	}
}
