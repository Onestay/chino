package commands

import (
	"github.com/onestay/chino/framework"
)

var Ping = framework.Command{
	Name:        "ping",
	Description: "Ping pong ping pong",
	Usage:       "!ping",
	Run: func(ctx framework.Context) {
		ctx.Send("pong")
	},
	Validator: func(ctx framework.Context) bool {
		return true
	},
}
