package models

import (
	"time"
)

type Championship struct {
	Base
	Name       string    `gorm:"size:100;not null" json:"name"`
	StartDate  time.Time `gorm:"not null" json:"start_date"`
	EndDate    time.Time `gorm:"not null" json:"end_date"`
	IsActive   bool      `gorm:"default:true" json:"is_active"`
	AthleticID int       `gorm:"not null" json:"athletic_id"`
	Athletic   Athletic  `gorm:"foreignKey:AthleticID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"athletic"`
}
