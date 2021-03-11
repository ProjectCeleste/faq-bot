package cmd

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestTriggerStatus(t *testing.T) {
	c := NewStatus()
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "server up?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "server up ?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "is the server up?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "the server is up?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "server down?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "server down ?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "is the server down?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "the server is down?"}}))

	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "server up"}}))
	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "server up"}}))
	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "is the server up"}}))
	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "the server is up"}}))

	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "look up the server?"}}))
	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "down server?"}}))
	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "are you on the server?"}}))
	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "i'm off, see you later on the server?"}}))
	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "i'm leaving the server? no that's off"}}))
}
