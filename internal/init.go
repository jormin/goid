package internal

import (
	"gitlab.wcxst.com/jormin/goid/internal/config"
	"gitlab.wcxst.com/jormin/golog/log"
)

func Init() {
	// 初始化配置
	config.Init()
	// 初始化日志
	log.Init(&config.LogCfg)
}
