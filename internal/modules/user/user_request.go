package user

import "time"

type CreateUserRequest struct {
	Name          string     `validate:"required,min=2,max=100" json:"name"`
	Email         string     `validate:"required,email,max=100" json:"email"`
	Password      string     `validate:"required,min=6,max=100" json:"password"`
	Telephone     string     `validate:"required,max=20" json:"telephone"`
	UniversityId  *int       `json:"university_id,omitempty"`
	StudentStatus string     `validate:"required,oneof=active graduated viewer prospective" json:"student_status"`
	DateOfBirth   *time.Time `json:"date_of_birth,omitempty"`
	AthleticRole  *string    `validate:"omitempty,oneof=member director president" json:"athletic_role,omitempty"`
}

type UpdateUserRequest struct {
	Name          string     `validate:"required,min=2,max=100" json:"name"`
	Email         string     `validate:"required,email,max=100" json:"email"`
	Telephone     string     `validate:"required,max=20" json:"telephone"`
	UniversityId  *int       `json:"university_id,omitempty"`
	StudentStatus string     `validate:"required,oneof=active graduated viewer prospective" json:"student_status"`
	DateOfBirth   *time.Time `json:"date_of_birth,omitempty"`
	AthleticRole  *string    `validate:"omitempty,oneof=member director president" json:"athletic_role,omitempty"`
	IsActive      bool       `json:"is_active,omitempty"`
}
