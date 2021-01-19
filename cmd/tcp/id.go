package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/snowflake"
	"gitlab.wcxst.com/jormin/go-tools/log"
	"gitlab.wcxst.com/jormin/go-tools/rpc"
	"gitlab.wcxst.com/jormin/goid/internal/config"
	"gitlab.wcxst.com/jormin/goid/internal/service"
)

// 初始化
func init() {
	log.SetPrefix("ID")
	home, _ := os.UserHomeDir()
	log.SetLogPath(fmt.Sprintf("%s/logs/id", home))
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
		log.Error("Start server on %s error: %v", addr, err)
		return
	}
	log.Info("Start server on %s success", addr)
}
