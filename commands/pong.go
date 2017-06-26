package commands

import (
"github.com/onestay/chino/framework"
)

var Pong = framework.Command{
	Name:        "pong",
	Description: "Ping pong ping pong",
	Usage:       "!pong",
	Run: func(ctx framework.Context) {
		ctx.Send("ping")
	},
	Validator: func(ctx framework.Context) bool {
		return true
	},
}
