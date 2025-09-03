package models

import (
	"time"
)

type Match struct {
	Base
	TournamentID int        `json:"tournament_id"`
	Tournament   Tournament `gorm:"foreignKey:TournamentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TeamAID      int        `json:"team_a_id"`
	TeamA        Team       `gorm:"foreignKey:TeamAID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TeamBID      int        `json:"team_b_id"`
	TeamB        Team       `gorm:"foreignKey:TeamBID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SportID      *uint      `json:"sport_id,omitempty"` // FK -> Sport (opcional)
	Sport        *Sport     `gorm:"foreignKey:SportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"sport,omitempty"`
	Location     string     `gorm:"size:200" json:"location"`
	DateTime     time.Time  `json:"date_time"`
}
