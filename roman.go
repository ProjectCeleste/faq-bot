package main

import (
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/martian/log"
	"github.com/sajari/fuzzy"
)

type roman struct {
	model       *fuzzy.Model
	httpClient  *http.Client
	caesarEmoji *discordgo.Emoji
}

func newRoman() *roman {
	return &roman{
		model: makeModel([]string{"roman", "out", "released"}),
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (q *roman) trigger(message []string) bool {
	romanFound := false
	releasedFound := false

	for _, w := range message {
		spell := q.model.SpellCheck(w)
		if spell == "roman" {
			romanFound = true
		} else if spell == "out" || spell == "released" {
			releasedFound = true
		}

		if romanFound && releasedFound {
			return true
		}
	}

	return false
}

func (q *roman) answer(s *discordgo.Session, m *discordgo.MessageCreate) {

	req, err := http.NewRequest(http.MethodGet, "https://api.projectceleste.com/gamedb/equipments?civilization=Roman", nil)
	if err != nil {
		log.Errorf("Could not create roman HTTP request: ", err.Error())
		return
	}

	status := ""
	req.Header.Set("Cache-Control", "no-cache")
	resp, err := q.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		status = "The romans are out! :partying_face:"
	} else if resp.StatusCode == http.StatusNotFound {
		status = "The romans are not released yet. "

		if q.caesarEmoji == nil {
			emoji, err := findEmoji(s, m.GuildID, "Caesar")
			if err != nil {
				log.Errorf(err.Error())
			}
			q.caesarEmoji = emoji
		}

		if q.caesarEmoji != nil {
			status += q.caesarEmoji.MessageFormat()
		}
	} else {
		return
	}

	message := &discordgo.MessageSend{
		Content:   status,
		Reference: m.Reference(),
	}
	s.ChannelMessageSendComplex(m.ChannelID, message)
}
