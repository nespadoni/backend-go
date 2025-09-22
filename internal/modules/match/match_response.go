package match

import (
	"backend-go/internal/models"
	"time"
)

type Response struct {
	ID           uint              `json:"id"`
	TournamentID int               `json:"tournament_id"`
	Tournament   models.Tournament `json:"tournament"`
	TeamAID      int               `json:"team_a_id"`
	TeamA        models.Team       `json:"team_a"`
	TeamBID      int               `json:"team_b_id"`
	TeamB        models.Team       `json:"team_b"`
	SportID      *uint             `json:"sport_id,omitempty"`
	Sport        *models.Sport     `json:"sport,omitempty"`
	Location     string            `json:"location"`
	DateTime     time.Time         `json:"date_time"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}

type ListResponse struct {
	ID           uint              `json:"id"`
	TournamentID int               `json:"tournament_id"`
	Tournament   TournamentSummary `json:"tournament"`
	TeamAID      int               `json:"team_a_id"`
	TeamA        TeamSummary       `json:"team_a"`
	TeamBID      int               `json:"team_b_id"`
	TeamB        TeamSummary       `json:"team_b"`
	SportID      *uint             `json:"sport_id,omitempty"`
	Sport        *SportSummary     `json:"sport,omitempty"`
	Location     string            `json:"location"`
	DateTime     time.Time         `json:"date_time"`
}

type TournamentSummary struct {
	ID           uint                `json:"id"`
	Name         string              `json:"name"`
	StartDate    time.Time           `json:"start_date"`
	EndDate      time.Time           `json:"end_date"`
	Championship ChampionshipSummary `json:"championship"`
}

type TeamSummary struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Logo     string `json:"logo,omitempty"`
	Athletic struct {
		ID         uint   `json:"id"`
		Name       string `json:"name"`
		University struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
		} `json:"university"`
	} `json:"athletic"`
}

type ChampionshipSummary struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type SportSummary struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
