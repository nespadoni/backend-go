package result

import (
	"backend-go/internal/models"
	"time"
)

type Response struct {
	Id         uint         `json:"id"`
	MatchID    int          `json:"match_id"`
	Match      models.Match `json:"match,omitempty"`
	TeamAScore int          `json:"team_a_score"`
	TeamBScore int          `json:"team_b_score"`
	Status     string       `json:"status"`
	IsLive     bool         `json:"is_live"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

type ListResponse struct {
	Id         uint   `json:"id"`
	MatchID    int    `json:"match_id"`
	TeamAScore int    `json:"team_a_score"`
	TeamBScore int    `json:"team_b_score"`
	Status     string `json:"status"`
	IsLive     bool   `json:"is_live"`
}

type ResultSummary struct {
	Id         uint   `json:"id"`
	MatchID    int    `json:"match_id"`
	TeamAScore int    `json:"team_a_score"`
	TeamBScore int    `json:"team_b_score"`
	Status     string `json:"status"`
}
