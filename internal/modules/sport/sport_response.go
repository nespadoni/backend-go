package sport

import (
	"backend-go/internal/models"
	"time"
)

type Response struct {
	Id           uint              `json:"id"`
	Name         string            `json:"name"`
	Description  string            `json:"description,omitempty"`
	Abbreviation string            `json:"abbreviation,omitempty"`
	Icon         string            `json:"icon,omitempty"`
	Category     string            `json:"category,omitempty"`
	MinPlayers   int               `json:"min_players"`
	MaxPlayers   int               `json:"max_players"`
	Rules        string            `json:"rules,omitempty"`
	IsActive     bool              `json:"is_active"`
	IsPopular    bool              `json:"is_popular"`
	Positions    []models.Position `json:"positions,omitempty"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}

type ListResponse struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	Abbreviation string `json:"abbreviation,omitempty"`
	Icon         string `json:"icon,omitempty"`
	Category     string `json:"category,omitempty"`
	MinPlayers   int    `json:"min_players"`
	MaxPlayers   int    `json:"max_players"`
	IsActive     bool   `json:"is_active"`
	IsPopular    bool   `json:"is_popular"`
}

type SportSummary struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation,omitempty"`
	Icon         string `json:"icon,omitempty"`
	Category     string `json:"category,omitempty"`
}
