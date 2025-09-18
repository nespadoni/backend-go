package tournament

import (
	"backend-go/internal/models"
	"time"
)

type Response struct {
	ID             uint                `json:"id"`
	Name           string              `json:"name"`
	StartDate      time.Time           `json:"start_date"`
	EndDate        time.Time           `json:"end_date"`
	ChampionshipID uint                `json:"championship_id"`
	Championship   models.Championship `json:"championship"`
	SportID        uint                `json:"sport_id"`
	Sport          models.Sport        `json:"sport"`
	CreatedAt      time.Time           `json:"created_at"`
	UpdatedAt      time.Time           `json:"updated_at"`
}

type ListResponse struct {
	ID             uint                `json:"id"`
	Name           string              `json:"name"`
	StartDate      time.Time           `json:"start_date"`
	EndDate        time.Time           `json:"end_date"`
	ChampionshipID uint                `json:"championship_id"`
	Championship   ChampionshipSummary `json:"championship"`
	SportID        uint                `json:"sport_id"`
	Sport          SportSummary        `json:"sport"`
}

type ChampionshipSummary struct {
	ID       uint            `json:"id"`
	Name     string          `json:"name"`
	IsActive bool            `json:"is_active"`
	Athletic AthleticSummary `json:"athletic"`
}

type AthleticSummary struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Logo       string `json:"logo,omitempty"`
	University struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"university"`
}

type SportSummary struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Modality string `json:"modality"`
}
