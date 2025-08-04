package result

import (
	"backend-go/internal/base"
)

type Result struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	MatchID uint `json:"match_id"`
	ScoreA  int  `json:"score_a"`
	ScoreB  int  `json:"score_b"`
	base.Base
}
