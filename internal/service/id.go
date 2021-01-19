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
func (id *ID) NewID(args *Args, res *int64) error {
	*res = id.Node.Generate().Int64()
	return nil
}
