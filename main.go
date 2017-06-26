package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/onestay/chino/commands"
	"github.com/onestay/chino/framework"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var token string
var cmdHandler *framework.CommandHandler
var prefix = "c>"

func init() {
	flag.StringVar(&token, "t", "", "Token of your bot")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.AddHandler(guildCreate)

	err = dg.Open()
	if err != nil {
		log.Fatalln("Error opening ws connection, ", err)
	}

	defer dg.Close()

	cmdHandler = framework.NewCommandHandlerWithDB("chino", "guilds")
	registerCommands()

	fmt.Println("Bot running. Press CTRL+C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//if len(m.Content) > len(prefix) && string(m.Content[0:len(prefix)]) != prefix {
	//	return
	//}
	//
	//if s.State.User.ID == m.Author.ID {
	//	return
	//}
	//command := strings.Split(m.Content, " ")[0]
	//command = strings.TrimLeft(command, prefix)

	go cmdHandler.OnMessage(m.Content, s, m)
}

func guildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	go cmdHandler.MS.CreateGuildConfig(g.ID)
}

func registerCommands() {
	go cmdHandler.AddCommand(commands.Ping)
	go cmdHandler.AddCommand(commands.Pong)
}
