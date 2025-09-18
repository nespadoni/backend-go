package tournament

import (
	"time"
)

type CreateRequest struct {
	Name           string    `validate:"required,min=3,max=100" json:"name"`
	StartDate      time.Time `validate:"required" json:"start_date"`
	EndDate        time.Time `validate:"required" json:"end_date"`
	ChampionshipID uint      `validate:"required,min=1" json:"championship_id"`
	SportID        uint      `validate:"required,min=1" json:"sport_id"`
}

type UpdateRequest struct {
	Name      string    `validate:"required,min=3,max=100" json:"name"`
	StartDate time.Time `validate:"required" json:"start_date"`
	EndDate   time.Time `validate:"required" json:"end_date"`
}
