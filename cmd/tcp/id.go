package main

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
	"gitlab.wcxst.com/jormin/goid/internal/config"
	"gitlab.wcxst.com/jormin/goid/internal/service"
	"gitlab.wcxst.com/jormin/golog/log"
	"gitlab.wcxst.com/jormin/gorpc/rpc"
)

// 初始化
func init() {
	log.SetPrefix("ID")
	log.SetLogPath("/var/log/goid")
}

func main() {
	addr := fmt.Sprintf(":%d", config.IP)
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal("get snowflake node error: %v", err)
	}
	err = rpc.NewServer(
		addr, &service.ID{Node: node},
	)
	if err != nil {
		log.Error("start server on %s error: %v", addr, err)
		return
	}
	log.Info("start server on %s success", addr)
}
