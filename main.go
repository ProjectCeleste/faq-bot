package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token, ok := os.LookupEnv("TOKEN")
	if !ok {
		log.Fatalf("Missing TOKEN env variable")
	}
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(readMessage)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Println("FAQ-Bot is running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	// Cleanly close down the Discord session.
	if err := dg.Close(); err != nil {
		log.Fatal(err)
	}
}

func readMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	words := strings.Fields(m.Content)
	for _, q := range questions {
		if q.trigger(words) {
			q.answer(s, m)
			break
		}
	}
}
