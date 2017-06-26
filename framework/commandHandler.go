package framework

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"fmt"
)

type Command struct {
	Name        string
	Description string
	Usage       string
	Run         func(ctx Context)
	Validator   func(ctx Context) bool
	GuildOny    bool
}

type CmdMap map[string]Command

type CommandHandler struct {
	Command CmdMap
	MS      *MongoSession
}

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make(CmdMap), nil}
}

func NewCommandHandlerWithDB(db, collection string) *CommandHandler {
	ms, err := InitDB(db, collection)
	if err != nil {
		panic(err)
	}

	return &CommandHandler{make(CmdMap), ms}
}

func (ch CommandHandler) AddCommand(c Command) {
	ch.Command[c.Name] = c
}

func (ch CommandHandler) OnMessage(message string, s *discordgo.Session, m *discordgo.MessageCreate) {
	// TODO replace hardcoded prefix with the prefix defined in main and check for a custom guild prefix
	prefix := "c>"

	// if the bot is the author; do nothing
	if s.State.User.ID == m.Author.ID {
		return
	}

	// this will check if the name of the command exists if a message
	// this is done because we don't want to call out to the db to check for custom prefix on every message which we would have to do
	hasCommand := make(chan bool)
	for key := range ch.Command {
		go func(f chan bool, commandName string) {
			if strings.Contains(message, commandName) {
				f <- true
			}
		}(hasCommand, key)
	}

	if !<-hasCommand {
		return
	}

	// this checks if message starts with a prefix
	if len(m.Content) > len(prefix) && string(m.Content[0:len(prefix)]) != prefix {
		return
	}


	command := strings.Split(m.Content, " ")[0]
	command = strings.TrimLeft(command, prefix)
	fmt.Println("Has command")

	c, found := ch.Command[command]
	if !found {
		return
	}

	args := strings.Split(m.Content, " ")
	args = append(args[:0], args[1:]...)

	var ctx Context

	if ch.MS != nil {
		ctx = *createContext(s, m, args, ch.MS)
	} else {
		ctx = *createContext(s, m, args, nil)
	}

	if !c.Validator(ctx) {
		ctx.Send("You don't have permission to run that command.")
		return
	}

	c.Run(ctx)
}
