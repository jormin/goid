package id

import (
	"testing"

	"gitlab.wcxst.com/jormin/golog/log"
)

// 测试生成ID
func TestClientNewID(t *testing.T) {
	c, err := NewClient("127.0.0.1:8888")
	// c, err := NewClient("111.229.255.133:31307")
	if err != nil {
		t.Errorf("get id client error: %v", err)
	}
	b, err := c.NewID()
	if err != nil {
		t.Errorf("get id error: %v", err)
	}
	log.Info("get id success: %d", b)
}
