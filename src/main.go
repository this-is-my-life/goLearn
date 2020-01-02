package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New("Bot " + os.Getenv("goToken"))
	if err != nil {
		fmt.Println(err)
		return
	}

	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println(event.User.Username, "is online!")
	s.UpdateStatus(0, "Testing discordgo")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	ori := m.Content
	if strings.HasPrefix(ori, "go>") {
		ori = strings.Replace(ori, "go>", "", 1)
		spOri := strings.Split(ori, " ")
		cmd := spOri[0]

		switch cmd {
		case "ping":
			s.ChannelMessageSend(m.ChannelID, "Pong!")
			break

		case "embed":
			emb := discordgo.MessageEmbed{Title: "Hello, world!", Description: "This Is Example Message for Embed"}
			s.ChannelMessageSendEmbed(m.ChannelID, &emb)
			break
		}
	}
}
