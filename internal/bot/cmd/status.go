package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ProjectCeleste/faq-bot/internal/detect"
	"github.com/bwmarrin/discordgo"
)

// Status say game server status when asked.
type Status struct {
	SentenceDetection
	httpClient *http.Client
}

type health struct {
	IsHealthy    bool `json:"isHealthy"`
	ResponseTime int  `json:"responseTime"`
}

// NewStatus create a new Thanks command.
func NewStatus() *Status {
	return &Status{
		SentenceDetection: SentenceDetection{
			Detector: &detect.SentenceDetector{
				Question: detect.QuestionOnly,
				Groups: []detect.WordGroup{
					{Words: []string{"server"}},
					{Words: []string{"up", "down", "on", "off"}},
				},
			},
		},
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

// Answer the game server status.
func (c *Status) Answer(s *discordgo.Session, m *discordgo.MessageCreate) error {
	req, err := http.NewRequest(http.MethodGet, "https://api.projectceleste.com/status/health", nil)
	if err != nil {
		sendErrorMessage(s, m)
		return fmt.Errorf("Could not create health HTTP request: %w", err)
	}

	status := ""
	req.Header.Set("Cache-Control", "no-cache")
	resp, err := c.httpClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		status = "UNREACHABLE :orange_circle:"
	} else {

		defer resp.Body.Close()
		health := health{}
		if err := json.NewDecoder(resp.Body).Decode(&health); err != nil {
			sendErrorMessage(s, m)
			return fmt.Errorf("Health response decoding error: %w", err)
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
	_, err = s.ChannelMessageSendComplex(m.ChannelID, message)
	return err
}
