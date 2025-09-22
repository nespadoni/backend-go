package playerstats

import (
	"backend-go/internal/models"
	"time"
)

type Response struct {
	Id            uint          `json:"id"`
	PlayerID      uint          `json:"player_id"`
	Player        models.Player `json:"player,omitempty"`
	MatchID       uint          `json:"match_id"`
	Match         models.Match  `json:"match,omitempty"`
	Goals         int           `json:"goals"`
	Assists       int           `json:"assists"`
	YellowCards   int           `json:"yellow_cards"`
	RedCards      int           `json:"red_cards"`
	MinutesPlayed int           `json:"minutes_played"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

type ListResponse struct {
	Id            uint `json:"id"`
	PlayerID      uint `json:"player_id"`
	MatchID       uint `json:"match_id"`
	Goals         int  `json:"goals"`
	Assists       int  `json:"assists"`
	YellowCards   int  `json:"yellow_cards"`
	RedCards      int  `json:"red_cards"`
	MinutesPlayed int  `json:"minutes_played"`
}

type PlayerStatsSummary struct {
	Id       uint `json:"id"`
	PlayerID uint `json:"player_id"`
	MatchID  uint `json:"match_id"`
	Goals    int  `json:"goals"`
	Assists  int  `json:"assists"`
}
