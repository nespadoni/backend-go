package models

import "time"

type Result struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	MatchID   uint      `json:"match_id"`
	ScoreA    int       `json:"score_a"`
	ScoreB    int       `json:"score_b"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
