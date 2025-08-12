package models

type Result struct {
	Base
	MatchID    int    `json:"match_id"`
	Match      Match  `gorm:"foreignKey:MatchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TeamAScore int    `gorm:"default:0" json:"team_a_score"`
	TeamBScore int    `gorm:"default:0" json:"team_b_score"`
	Status     string `gorm:"size:20;default:'scheduled'" json:"status"` // scheduled, live, finished, postponed
	IsLive     bool   `gorm:"default:false" json:"is_live"`
}
