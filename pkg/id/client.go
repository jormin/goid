package id

import (
	rpc2 "net/rpc"

	"gitlab.wcxst.com/jormin/go-tools/rpc"
	"gitlab.wcxst.com/jormin/goid/internal/config"
	"gitlab.wcxst.com/jormin/goid/internal/service"
)

// 客户端
type Client struct {
	RpcClient *rpc2.Client
}

// 采集
func (c *Client) NewID() ([]byte, error) {
	var res []byte
	err := c.RpcClient.Call(
		config.IDCall, &service.Args{}, &res,
	)
	return res, err
}

// 新客户端连接
func NewClient(addr string) (*Client, error) {
	rc, err := rpc.NewClient(addr)
	if err != nil {
		return nil, err
	}
	client := Client{
		RpcClient: rc,
	}
	return &client, nil
}
