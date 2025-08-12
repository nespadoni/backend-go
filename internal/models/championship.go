package models

import (
	"time"
)

type Championship struct {
	Base
	Name      string    `gorm:"size:100;not null" json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
