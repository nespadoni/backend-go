package championship

import (
	"backend-go/internal/base"
	"backend-go/internal/match"
	"time"
)

type Championship struct {
	ID        uint          `gorm:"primaryKey" json:"id"`
	Name      string        `gorm:"size:100;not null" json:"name"`
	SportID   uint          `json:"sport_id"`
	StartDate time.Time     `json:"start_date"`
	EndDate   time.Time     `json:"end_date"`
	Matches   []match.Match `gorm:"foreignKey:ChampionshipID" json:"matches"`
	base.Base
}
