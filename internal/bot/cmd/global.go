package cmd

import (
	"time"

	"github.com/ProjectCeleste/faq-bot/internal/detect"
	"github.com/bwmarrin/discordgo"
)

// GlobalCommand a simple generic command that can be triggered by any user.
// These commands trigger only once every Timeout.
type GlobalCommand struct {
	TimeoutTrigger
	Response string
}

// Answer the message defined in this GlobalCommand.
func (c *GlobalCommand) Answer(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := &discordgo.MessageSend{
		Content:   c.Response,
		Reference: m.Reference(),
	}
	_, err := s.ChannelMessageSendComplex(m.ChannelID, message)
	return err
}

// NewGlobalCommand create a new basic targeted command with a default Timeout of 2 hours.
func NewGlobalCommand(keywords [][]string, response string) *GlobalCommand {
	groups := []detect.WordGroup{}
	for _, g := range keywords {
		groups = append(groups, detect.WordGroup{Words: g})
	}
	return &GlobalCommand{
		TimeoutTrigger: TimeoutTrigger{
			CommandTrigger: &SentenceDetection{
				Detector: &detect.SentenceDetector{
					Question: detect.QuestionBoth,
					Groups:   groups,
				},
			},
			Timeout: time.Hour * 24,
		},
		Response: response,
	}
}

// CreateGlobalCommands create a slice of pre-built global commands.
func CreateGlobalCommands() []Command {
	return []Command{
		NewGlobalCommand([][]string{{"clinicus"}}, "The Clinicus may appear at first a doddering old man, but inside he is quite deft. It is always easy to begin a war, but very difficult to stop one."),
		NewGlobalCommand([][]string{{"legionary"}}, "I have always been of the opinion that unpopularity earned by doing what is right is not unpopularity at all, but glory."),
		NewGlobalCommand([][]string{{"centurion"}}, "Great empires are not maintained by timidity."),
		NewGlobalCommand([][]string{{"scorpion"}}, "For the last time, Scorpios are Roman anti-Infantry siege craft. Scorpions are what crawls out of the mouths of dead Egyptians."),
		NewGlobalCommand([][]string{{"decurion"}}, "Roma Invicta! We gather strength as we go!"),
		NewGlobalCommand([][]string{{"aquilifer"}}, "Where there is unity, there is victory. We must defend our Eagle with our lives!"),
		NewGlobalCommand([][]string{{"gold"}}, "The sinews of war are infinite money."),
	}
}
