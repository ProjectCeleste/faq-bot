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

// SentenceDetection a generic command triggered by a SentenceDetector.
type SentenceDetection struct {
	Detector *detect.SentenceDetector
}

// Trigger returns true if the message matches the trigger criteria for this command.
func (s *SentenceDetection) Trigger(message *discordgo.MessageCreate) bool {
	sentence := strings.ToLower(strings.TrimSpace(message.Content))
	sentence = quotes.ReplaceAllString(sentence, "")
	return s.Detector.Detect(sentence)
}

// TargetedSentenceDetection a SentenceDetection that only triggers if the message
// is from the user defined in TargetUserID.
type TargetedSentenceDetection struct {
	SentenceDetection
	TargetUserID string
}

// Trigger returns true if the message's author is this command's target and the message matches
// the trigger criteria.
func (s *TargetedSentenceDetection) Trigger(message *discordgo.MessageCreate) bool {
	if message.Author.ID == s.TargetUserID {
		return s.SentenceDetection.Trigger(message)
	}
	return false
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
