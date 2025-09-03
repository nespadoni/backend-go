package models

import (
	"time"
)

type Tournament struct {
	Base
	Name           string       `gorm:"size:100;not null" json:"name"`
	ChampionshipID uint         `json:"championship_id"` // Mudança: int -> uint
	Championship   Championship `gorm:"foreignKey:ChampionshipID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SportID        uint         `json:"sport_id"` // Mudança: int -> uint
	Sport          Sport        `gorm:"foreignKey:SportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StartDate      time.Time    `json:"start_date"`
	EndDate        time.Time    `json:"end_date"`
}
