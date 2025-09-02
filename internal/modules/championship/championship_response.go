package championship

import (
	"backend-go/internal/models"
	"time"
)

type Response struct {
	Id        uint            `json:"id"`
	Name      string          `json:"name"`
	StartDate time.Time       `json:"start_date"`
	EndDate   time.Time       `json:"end_date"`
	IsActive  bool            `json:"is_active"`
	Athletic  models.Athletic `json:"athletic"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type ListResponse struct {
	Id        uint            `json:"id"`
	Name      string          `json:"name"`
	StartDate time.Time       `json:"start_date"`
	EndDate   time.Time       `json:"end_date"`
	IsActive  bool            `json:"is_active"`
	Athletic  AthleticSummary `json:"athletic"`
}

type AthleticSummary struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Logo       string `json:"logo,omitempty"`
	University struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"university"`
}
