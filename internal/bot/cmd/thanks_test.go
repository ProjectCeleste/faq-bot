package cmd

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestTriggerThanks(t *testing.T) {
	c := NewThanks()
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "thanks crassus"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "thank you crassus"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "ty crassus"}}))
	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "crassus ty"}}))
	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "thanks, crassus?"}}))
	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "thank you? crassus?"}}))
	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "no thanks crassus"}}))
}
