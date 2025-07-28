package models

type Result struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	MatchID uint `json:"match_id"`
	ScoreA  int  `json:"score_a"`
	ScoreB  int  `json:"score_b"`
	Base
}
