package cmd

import (
	"testing"
	"time"

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
	assert.False(t, c.Trigger(&discordgo.MessageCreate{
		Message: &discordgo.Message{
			Author: &discordgo.User{
				ID: systemGlitch,
			},
			Content: "The Romans are capable of many things",
		},
	}))

	c = NewTargetedCommand(systemGlitch, [][]string{{"awesome"}}, "Not as awesome as you are, master! I am looking forward to our next lesson.")
	assert.True(t, c.Trigger(&discordgo.MessageCreate{
		Message: &discordgo.Message{
			Author: &discordgo.User{
				ID: systemGlitch,
			},
			Content: "Crassus is. awesome.",
		},
	}))
	c.lastResponse = time.Time{}
	assert.True(t, c.Trigger(&discordgo.MessageCreate{
		Message: &discordgo.Message{
			Author: &discordgo.User{
				ID: systemGlitch,
			},
			Content: "Crassus is awesome, isn't it?",
		},
	}))
}
