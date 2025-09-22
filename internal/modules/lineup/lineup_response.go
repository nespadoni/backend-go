package lineup

import (
	"backend-go/internal/models"
	"time"
)

type Response struct {
	Id        uint          `json:"id"`
	MatchID   uint          `json:"match_id"`
	Match     models.Match  `json:"match,omitempty"`
	TeamID    uint          `json:"team_id"`
	PlayerID  uint          `json:"player_id"`
	Player    models.Player `json:"player,omitempty"`
	Position  string        `json:"position"`
	IsStarter bool          `json:"is_starter"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type ListResponse struct {
	Id        uint   `json:"id"`
	MatchID   uint   `json:"match_id"`
	TeamID    uint   `json:"team_id"`
	PlayerID  uint   `json:"player_id"`
	Position  string `json:"position"`
	IsStarter bool   `json:"is_starter"`
}

type LineupSummary struct {
	Id       uint   `json:"id"`
	MatchID  uint   `json:"match_id"`
	PlayerID uint   `json:"player_id"`
	Position string `json:"position"`
}
