package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sajari/fuzzy"
)

type thanks struct {
	model *fuzzy.Model
}

func newThanks() *thanks {
	return &thanks{
		model: makeModel([]string{"thank", "thanks", "you", "crassus"}),
	}
}

func (q *thanks) trigger(message []string) bool {
	thanksFound := false
	thankFound := false
	youFound := false
	crassusFound := false

	for _, w := range message {
		spell := q.model.SpellCheck(w)
		if spell == "thanks" {
			thanksFound = true
		} else if spell == "thank" {
			thankFound = true
		} else if spell == "you" {
			youFound = true
		} else if spell == "crassus" {
			crassusFound = true
		}

		if (thanksFound || (thankFound && youFound)) && crassusFound {
			return true
		}
	}

	return false
}

func (q *thanks) answer(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := &discordgo.MessageSend{
		Content:   "You are very welcome! I will see you soon inside the gates of the Eternal City of Rome!",
		Reference: m.Reference(),
	}
	s.ChannelMessageSendComplex(m.ChannelID, message)
}
