package service

import (
	"github.com/bwmarrin/snowflake"
)

type ID struct {
	Node *snowflake.Node
}

type Args struct {
}

// Generate a snowflake ID.
func (id *ID) NewID(args *Args, res *[]byte) error {
	*res = id.Node.Generate().Bytes()
	return nil
}
