package auth

type RegisterRequest struct {
	Name         string  `json:"name" validate:"required,min=2,max=100"`
	Email        string  `json:"email" validate:"required,email,max=100"`
	Password     string  `json:"password" validate:"required,min=6,max=100"`
	Telephone    string  `json:"telephone" validate:"required,max=20"`
	UniversityID *string `json:"universityId,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
