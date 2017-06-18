package framework

import "github.com/bwmarrin/discordgo"

type Context struct {
	S *discordgo.Session
	M *discordgo.MessageCreate
	Args []string
}

func (ctx Context) Send(message string) {
	ctx.S.ChannelMessageSend(ctx.M.ChannelID, message)
}

func createContext(s *discordgo.Session, m *discordgo.MessageCreate, args []string) *Context {
	return &Context{s, m, args}
}