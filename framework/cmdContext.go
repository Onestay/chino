package framework

import "github.com/bwmarrin/discordgo"

type Context struct {
	S *discordgo.Session
	M *discordgo.MessageCreate
	Args []string
	MS *MongoSession
}

func createContext(s *discordgo.Session, m *discordgo.MessageCreate, args []string, ms *MongoSession) *Context {
	return &Context{s, m, args, ms}
}

// helper functions provided by the ctx object

func (ctx Context) Send(message string) {
	ctx.S.ChannelMessageSend(ctx.M.ChannelID, message)
}