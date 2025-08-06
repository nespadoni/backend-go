package stats

import (
	"backend-go/internal/base"
	"backend-go/internal/match"
	"backend-go/internal/player"
)

type PlayerStats struct {
	base.Base
	PlayerID      uint          `json:"player_id"`
	Player        player.Player `gorm:"foreignKey:PlayerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MatchID       uint          `json:"match_id"`
	Match         match.Match   `gorm:"foreignKey:MatchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Goals         int           `gorm:"default:0" json:"goals"`
	Assists       int           `gorm:"default:0" json:"assists"`
	YellowCards   int           `gorm:"default:0" json:"yellow_cards"`
	RedCards      int           `gorm:"default:0" json:"red_cards"`
	MinutesPlayed int           `gorm:"default:0" json:"minutes_played"`
}
