package cmd

import (
	"log"

	"github.com/ProjectCeleste/faq-bot/internal/detect"
	"github.com/bwmarrin/discordgo"
)

// Romans say romans release date.
type Romans struct {
	SentenceDetection
	caesarEmoji *discordgo.Emoji
}

// NewRomans create a new Romans command.
func NewRomans() *Romans {
	return &Romans{
		SentenceDetection: SentenceDetection{
			Detector: &detect.SentenceDetector{
				Question: true,
				Groups: []detect.WordGroup{
					{Words: []string{"when", "are", "have", "has"}},
					{Words: []string{"romans", "roman"}},
					{Words: []string{"out", "ready", "released", "release", "coming"}},
				},
				Variants: []detect.SentenceDetector{
					{
						Question: true,
						Groups: []detect.WordGroup{
							{Words: []string{"when"}},
							{Words: []string{"play", "access"}},
							{Words: []string{"romans", "roman"}},
						},
					},
					{
						Question: true,
						Groups: []detect.WordGroup{
							{Words: []string{"romans", "roman"}},
							{Words: []string{"when", "out"}},
						},
					},
				},
			},
		},
	}
}

// Answer the romans release date.
func (c *Romans) Answer(s *discordgo.Session, m *discordgo.MessageCreate) error {
	status := "Beware the Ides of March! We shall throw open the gates of Rome on **March 15**. See you all in the Eternal City! "

	if c.caesarEmoji == nil {
		emoji, emojiErr := findEmoji(s, m.GuildID, "Crassus")
		if emojiErr != nil {
			log.Printf("Could not find emoji %q: %v\n", "Crassus", emojiErr)
		}
		c.caesarEmoji = emoji
	}

	if c.caesarEmoji != nil {
		status += c.caesarEmoji.MessageFormat()
	}

	message := &discordgo.MessageSend{
		Content:   status,
		Reference: m.Reference(),
	}
	_, err := s.ChannelMessageSendComplex(m.ChannelID, message)
	return err
}
