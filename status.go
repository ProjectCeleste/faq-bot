package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/martian/log"
	"github.com/sajari/fuzzy"
)

type serverStatus struct {
	model      *fuzzy.Model
	httpClient *http.Client
}

type health struct {
	IsHealthy    bool `json:"isHealthy"`
	ResponseTime int  `json:"responseTime"`
}

func newServerStatus() *serverStatus {
	return &serverStatus{
		model: makeModel([]string{"server", "down", "offline"}),
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (q *serverStatus) trigger(message []string) bool {
	serverFound := false
	downFound := false

	for _, w := range message {
		spell := q.model.SpellCheck(w)
		if spell == "server" {
			serverFound = true
		} else if spell == "down" || spell == "offline" {
			downFound = true
		}

		if serverFound && downFound && findQuestionMark(message) {
			return true
		}
	}

	return false
}

func (q *serverStatus) answer(s *discordgo.Session, m *discordgo.MessageCreate) {

	req, err := http.NewRequest(http.MethodGet, "https://api.projectceleste.com/status/health", nil)
	if err != nil {
		log.Errorf("Could not create health HTTP request: ", err.Error())
		return
	}

	status := ""
	req.Header.Set("Cache-Control", "no-cache")
	resp, err := q.httpClient.Do(req)
	if err != nil {
		status = "UNREACHABLE :orange_circle:"
	} else {
		if resp.StatusCode != http.StatusOK {
			return
		}

		defer resp.Body.Close()
		health := health{}
		if err := json.NewDecoder(resp.Body).Decode(&health); err != nil {
			log.Errorf("Health response: ", err.Error())
			return
		}

		if health.IsHealthy {
			status = "UP :green_circle:"
		} else {
			status = "DOWN :red_circle:"
		}
	}

	message := &discordgo.MessageSend{
		Content:   "**Game Server status**: " + status,
		Reference: m.Reference(),
	}
	s.ChannelMessageSendComplex(m.ChannelID, message)
}
