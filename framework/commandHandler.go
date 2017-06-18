package framework

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Command struct {
	Name        string
	Description string
	Usage       string
	Run         func(ctx Context)
	Validator   func(ctx Context) bool
}

type CmdMap map[string]Command

type CommandHandler struct {
	Command CmdMap
}

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make(CmdMap)}
}

func (ch CommandHandler) AddCommand(c Command) {
	ch.Command[c.Name] = c
}

func (ch CommandHandler) ExecuteCommand(name string, s *discordgo.Session, m *discordgo.MessageCreate) {
	c, found := ch.Command[name]
	if !found {
		return
	}

	args := strings.Split(m.Content, " ")
	args = append(args[:0], args[1:]...)
	ctx := *createContext(s, m, args)

	if !c.Validator(ctx) {
		ctx.Send("You don't have permission to run that command.")
		return
	}

	c.Run(ctx)
}
