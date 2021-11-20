package main

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
	"github.com/gofiber/fiber/v2"
)

var node *snowflake.Node

func init() {
	node, _ = snowflake.NewNode(1)
}

func main() {
	app := fiber.New()

	// 健康检测
	app.Get(
		"/health", func(c *fiber.Ctx) error {
			return c.SendString("ok")
		},
	)

	// ID
	app.Get(
		"/", func(c *fiber.Ctx) error {
			return c.SendString(fmt.Sprintf("%d", NewID()))
		},
	)

	_ = app.Listen(":80")
}

// NewID 生成ID
func NewID() int64 {
	return node.Generate().Int64()
}
