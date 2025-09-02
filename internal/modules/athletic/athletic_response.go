package athletic

import (
	"backend-go/internal/models"
	"time"
)

type Response struct {
	Id             uint              `json:"id"`
	Name           string            `json:"name"`
	Description    string            `json:"description"`
	Logo           string            `json:"logo,omitempty"`
	CoverImage     string            `json:"cover_image,omitempty"`
	IsActive       bool              `json:"is_active"`
	IsPublic       bool              `json:"is_public"`
	FollowersCount int               `json:"followers_count"`
	University     models.University `json:"university"`
	Creator        *models.User      `json:"creator,omitempty"`
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
}

type ListResponse struct {
	Id             uint              `json:"id"`
	Name           string            `json:"name"`
	Description    string            `json:"description"`
	Logo           string            `json:"logo,omitempty"`
	IsActive       bool              `json:"is_active"`
	IsPublic       bool              `json:"is_public"`
	FollowersCount int               `json:"followers_count"`
	University     models.University `json:"university"`
}
