package championship

import (
	"time"
)

type CreateRequest struct {
	Name       string    `validate:"required,min=3,max=100" json:"name"`
	StartDate  time.Time `validate:"required" json:"start_date"`
	EndDate    time.Time `validate:"required" json:"end_date"`
	AthleticId uint      `validate:"required,min=1" json:"athletic_id"`
}

type UpdateRequest struct {
	Name      string    `validate:"required,min=3,max=100" json:"name"`
	StartDate time.Time `validate:"required" json:"start_date"`
	EndDate   time.Time `validate:"required" json:"end_date"`
}

type UpdateStatusRequest struct {
	IsActive bool `json:"is_active"`
}
