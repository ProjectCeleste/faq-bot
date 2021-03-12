package cmd

import (
	"github.com/ProjectCeleste/faq-bot/internal/detect"
	"github.com/bwmarrin/discordgo"
)

// Thanks say "you're welcome" when the bot is thanked.
type Thanks struct {
	SentenceDetection
}

// NewThanks create a new Thanks command.
func NewThanks() *Thanks {
	return &Thanks{
		SentenceDetection: SentenceDetection{
			Detector: &detect.SentenceDetector{
				Question: false,
				Groups: []detect.WordGroup{
					{
						Negate: true,
						Words:  []string{"not", "no"},
					},
					{Words: []string{"thanks", "thank", "ty"}},
					{Words: []string{"crassus"}},
				},
			},
		},
	}
}

// Answer "you're welcome".
func (c *Thanks) Answer(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := &discordgo.MessageSend{
		Content:   "You are very welcome! I will see you soon inside the gates of the Eternal City of Rome!",
		Reference: m.Reference(),
	}
	_, err := s.ChannelMessageSendComplex(m.ChannelID, message)
	return err
}
