package bot

import (
	"fmt"

	"github.com/ProjectCeleste/faq-bot/internal/bot/cmd"
	"github.com/bwmarrin/discordgo"
)

// Bot a generic Discord bot.
type Bot struct {
	Commands []cmd.Command
	session  *discordgo.Session
}

// Connect create a Discord session authentified with the given token.
func (b *Bot) Connect(token string) error {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return fmt.Errorf("Error creating Discord session: %w", err)
	}

	session.AddHandler(b.readMessage)

	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	if err := session.Open(); err != nil {
		return fmt.Errorf("Error opening connection: %w", err)
	}

	b.session = session

	return nil
}

// Close the Discord session.
func (b *Bot) Close() error {
	return b.session.Close()
}

func (b *Bot) readMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	for _, c := range b.Commands {
		if c.Trigger(m) {
			c.Answer(s, m)
			break
		}
	}
}

// NewFAQBot create a new Celeste FAQ Bot with preloaded commands.
func NewFAQBot() *Bot {
	commands := []cmd.Command{
		cmd.NewStatus(),
		cmd.NewRomans(),
		cmd.NewThanks(),
		cmd.NewGameNotAvailable(),
		cmd.NewUpdateLauncher(),
		cmd.NewInstallLinux(),
		cmd.NewGameDoesntStart(),
		cmd.NewChatEmpty(),
		cmd.NewChangeUsername(),
	}
	targetedCmd := cmd.CreateTargetedCommands()
	commands = append(commands, targetedCmd...)
	globalCmd := cmd.CreateGlobalCommands()
	commands = append(commands, globalCmd...)
	bot := &Bot{
		Commands: commands,
	}
	return bot
}
