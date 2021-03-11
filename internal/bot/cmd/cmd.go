package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ProjectCeleste/faq-bot/internal/detect"
	"github.com/bwmarrin/discordgo"
)

var (
	quotes = regexp.MustCompile(`"(.*?)"`)
)

// Command a chatbot operation that has specific trigger conditions.
type Command interface {
	Trigger(message *discordgo.MessageCreate) bool
	Answer(s *discordgo.Session, m *discordgo.MessageCreate) error
}

// SentenceDetectionCommand a generic command triggered by a SentenceDetector.
type SentenceDetectionCommand struct {
	Detector *detect.SentenceDetector
}

// Trigger returns true if the message matches the trigger criteria for this command.
func (s *SentenceDetectionCommand) Trigger(message *discordgo.MessageCreate) bool {
	sentence := strings.ToLower(strings.TrimSpace(message.Content))
	sentence = quotes.ReplaceAllString(sentence, "")
	return s.Detector.Detect(sentence)
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

func sendErrorMessage(s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.Message, error) {
	message := &discordgo.MessageSend{
		Content:   ":worried: I am truly sorry, but I am not feeling very well right now. Please ask again later.",
		Reference: m.Reference(),
	}
	return s.ChannelMessageSendComplex(m.ChannelID, message)
}
