package service

import (
	"github.com/bwmarrin/snowflake"
	"gitlab.wcxst.com/jormin/golog/log"
)

type ID struct {
	Node *snowflake.Node
}

type Args struct {
}

// Generate a snowflake ID.
func (id *ID) NewID(args *Args, res *int64) error {
	*res = id.Node.Generate().Int64()
	log.Info("generate id %d", *res)
	return nil
}
