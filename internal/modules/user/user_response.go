package user

import (
	"backend-go/internal/models"
	"time"
)

type Response struct {
	Id              uint               `json:"id"`
	Name            string             `json:"name"`
	Email           string             `json:"email"`
	Telephone       string             `json:"telephone"`
	ProfilePhotoURL *string            `json:"profile_photo_url,omitempty"`
	University      *models.University `json:"university,omitempty"`
	StudentStatus   string             `json:"student_status"`
	DateOfBirth     *time.Time         `json:"date_of_birth,omitempty"`
	AthleticRole    *string            `json:"athletic_role,omitempty"`
	JoinDate        *time.Time         `json:"join_date,omitempty"`
	IsActive        bool               `json:"is_active"`
	CreatedAt       time.Time          `json:"created_at"`
}
