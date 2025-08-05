package match

import (
	"backend-go/internal/base"
	"backend-go/internal/tournament"
	"time"
)

type Match struct {
	base.Base
	TournmentID int                   `json:"tournment_id"`
	Tournment   tournament.Tournament `gorm:"foreignKey:TournmentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TeamAID     int                   `json:"team_a_id"`
	TeamBID     int                   `json:"team_b_id"`
	Location    string                `gorm:"size:200" json:"location"`
	DateTime    time.Time             `json:"date_time"`
}
