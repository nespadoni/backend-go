package user

type CreateUserRequest struct {
	Name         string `validate:"required" json:"name"`
	Email        string `validate:"required" json:"email"`
	Password     string `validate:"required" json:"password"`
	Telephone    string `validate:"required" json:"telephone"`
	UniversityId int    `validate:"required" json:"university_id"`
}

type UpdateUserRequest struct {
	Name         string `validate:"required" json:"name"`
	Email        string `validate:"required" json:"email"`
	Password     string `validate:"required" json:"password"`
	Telephone    string `validate:"required" json:"telephone"`
	UniversityId int    `validate:"required" json:"university_id"`
	RoleId       int    `validate:"required" json:"role_id"`
}
