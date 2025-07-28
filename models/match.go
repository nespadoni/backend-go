package models

import "time"

type Match struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	ChampionshipID uint      `json:"championship_id"`
	TeamAID        uint      `json:"team_a_id"`
	TeamBID        uint      `json:"team_b_id"`
	Location       string    `gorm:"size:200" json:"location"`
	DateTime       time.Time `json:"date_time"`
	Result         *Result   `gorm:"foreignKey:MatchID" json:"result,omitempty"`
	Base
}
