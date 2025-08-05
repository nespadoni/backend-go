package championship

import (
	"backend-go/internal/base"
	"time"
)

type Championship struct {
	base.Base
	Name      string    `gorm:"size:100;not null" json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
