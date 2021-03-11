package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ProjectCeleste/faq-bot/internal/bot"
)

const (
	exitMissingToken = 1
	exitConnectError = 2
	exitCloseError   = 3
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	errLogger := log.New(os.Stderr, "", log.LstdFlags)

	token, ok := os.LookupEnv("TOKEN")
	if !ok {
		errLogger.Println("Missing TOKEN env variable")
		os.Exit(exitMissingToken)
	}

	bot := bot.NewFAQBot()

	if err := bot.Connect(token); err != nil {
		errLogger.Println(err)
		os.Exit(exitConnectError)
	}

	// Wait here until CTRL-C or other term signal is received.
	logger.Println("FAQ-Bot is running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	// Cleanly close down the Discord session.
	if err := bot.Close(); err != nil {
		errLogger.Println(err)
		os.Exit(exitCloseError)
	}
}
