package models

import (
	"time"
)

type Championship struct {
	Base
	Name       string    `gorm:"size:100;not null" json:"name"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	AthleticID int       `json:"athletic_id"`                                                         // ADICIONAR
	Athletic   Athletic  `gorm:"foreignKey:AthleticID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // ADICIONAR
}
