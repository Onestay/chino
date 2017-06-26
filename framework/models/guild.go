package models

import (
	"gopkg.in/mgo.v2/bson"
)

type GuildConfig struct {
	ID               bson.ObjectId `bson:"_id,omitempty"`
	GuildID          string        `bson:"guild_id"`
	DisabledCommands []string      `bson:"disabled_commands"`
	Prefix           string        `bson:"prefix"`
}
