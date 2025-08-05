package result

import (
	"backend-go/internal/base"
	"backend-go/internal/match"
)

type Result struct {
	base.Base
	MatchID int         `json:"match_id"`
	Match   match.Match `gorm:"foreignKey:MatchID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ScoreA  int         `json:"score_a"`
	ScoreB  int         `json:"score_b"`
}
