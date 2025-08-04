package match

import (
	"backend-go/internal/base"
	"backend-go/internal/result"
	"time"
)

type Match struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	ChampionshipID uint           `json:"championship_id"`
	TeamAID        uint           `json:"team_a_id"`
	TeamBID        uint           `json:"team_b_id"`
	Location       string         `gorm:"size:200" json:"location"`
	DateTime       time.Time      `json:"date_time"`
	Result         *result.Result `gorm:"foreignKey:MatchID" json:"result,omitempty"`
	base.Base
}
