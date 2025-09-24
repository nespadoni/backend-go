package player

import (
	"backend-go/internal/models"
	"time"
)

type Response struct {
	Id         uint            `json:"id"`
	Name       string          `json:"name"`
	Nickname   *string         `json:"nickname"`
	TeamID     int             `json:"team_id"`
	Team       models.Team     `json:"team,omitempty"`
	PositionID int             `json:"position_id"`
	Position   models.Position `json:"position,omitempty"`
	UserID     *int            `json:"user_id,omitempty"`
	User       *models.User    `json:"user,omitempty"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}

type ListResponse struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	TeamID     int    `json:"team_id"`
	PositionID int    `json:"position_id"`
	UserID     *int   `json:"user_id,omitempty"`
}

type PlayerSummary struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	TeamID int    `json:"team_id"`
}
