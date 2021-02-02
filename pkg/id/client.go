package id

import (
	"fmt"
	rpc2 "net/rpc"

	"gitlab.wcxst.com/jormin/goid/internal/config"
	"gitlab.wcxst.com/jormin/goid/internal/service"
	"gitlab.wcxst.com/jormin/gorpc/rpc"
)

// 客户端
type Client struct {
	RpcClient *rpc2.Client
}

// 配置
type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// 采集
func (c *Client) NewID() (int64, error) {
	var res int64
	err := c.RpcClient.Call(
		config.IDCall, &service.Args{}, &res,
	)
	return res, err
}

// 新客户端连接
func NewClient(cfg *Config) (*Client, error) {
	rc, err := rpc.NewClient(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		return nil, err
	}
	client := Client{
		RpcClient: rc,
	}
	return &client, nil
}
