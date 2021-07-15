package id

import (
	"testing"

	"github.com/jormin/golog/log"
)

// 测试生成ID
func TestClientNewID(t *testing.T) {
	// c, err := NewClient("127.0.0.1:8888")
	c, err := NewClient(
		&Config{
			Host: "82.157.101.82",
			Port: 31339,
		},
	)
	if err != nil {
		t.Errorf("get id client error: %v", err)
	}
	b, err := c.NewID()
	if err != nil {
		t.Errorf("get id error: %v", err)
	}
	log.Info("get id success: %d", b)
}
