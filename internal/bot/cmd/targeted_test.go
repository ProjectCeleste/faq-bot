package cmd

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestTriggerTargeted(t *testing.T) {
	c := NewTargetedCommand(vick, [][]string{{"romans"}}, "Well Vick, aren’t you a beacon of ray and sunshine today. It is so nice to have you back.")
	assert.True(t, c.Trigger(&discordgo.MessageCreate{
		Message: &discordgo.Message{
			Author: &discordgo.User{
				ID: vick,
			},
			Content: "The Romans are capable of many things",
		},
	}))
}
