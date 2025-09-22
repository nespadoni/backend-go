package match

import (
	"time"
)

type CreateRequest struct {
	TournamentID int       `validate:"required,min=1" json:"tournament_id"`
	TeamAID      int       `validate:"required,min=1" json:"team_a_id"`
	TeamBID      int       `validate:"required,min=1" json:"team_b_id"`
	SportID      *uint     `validate:"omitempty,min=1" json:"sport_id,omitempty"`
	Location     string    `validate:"required,min=3,max=200" json:"location"`
	DateTime     time.Time `validate:"required" json:"date_time"`
}

type UpdateRequest struct {
	Location string    `validate:"required,min=3,max=200" json:"location"`
	DateTime time.Time `validate:"required" json:"date_time"`
}
