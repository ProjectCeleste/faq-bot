package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/sajari/fuzzy"
)

type question interface {
	trigger(message []string) bool
	answer(s *discordgo.Session, m *discordgo.MessageCreate)
}

var (
	questions = []question{
		newServerStatus(),
		newRoman(),
	}
)

func makeModel(words []string) *fuzzy.Model {
	model := fuzzy.NewModel()

	model.SetThreshold(2)

	model.Train(words)
	return model
}

func findEmoji(s *discordgo.Session, guildID, name string) (*discordgo.Emoji, error) {
	emojis, err := s.GuildEmojis(guildID)
	if err != nil {
		return nil, err
	}

	for _, e := range emojis {
		if e.Name == name {
			return e, nil
		}
	}

	return nil, fmt.Errorf("Emoji not found: %s", name)
}
