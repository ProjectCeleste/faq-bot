package cmd

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestTriggerRomans(t *testing.T) {
	c := NewRomans()
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "when will the romans come out?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "  When will the romans come out?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "when can i play the romans?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "are the romans ready?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "are the romans out?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "are the romans released?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "when will the romans release?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "are the romans out yet?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "when are romans being released?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "romans when?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "have the romans been released?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "romans out?"}}))
	assert.True(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "when are the romans coming?"}}))

	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "when will the ro  mans come out?"}}))
	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "when will the romans come out"}}))
	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "the romans are out. When are you going to play the romans?"}}))
	assert.False(t, c.Trigger(&discordgo.MessageCreate{Message: &discordgo.Message{Content: "When are you going to p l a y the romans? they release march 15?"}}))
}
