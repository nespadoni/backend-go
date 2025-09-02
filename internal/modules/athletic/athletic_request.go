package athletic

type CreateRequest struct {
	Name         string `validate:"required,min=3,max=100" json:"name"`
	Description  string `validate:"required,min=10,max=1000" json:"description"`
	Logo         string `validate:"omitempty,url,max=255" json:"logo,omitempty"`
	CoverImage   string `validate:"omitempty,url,max=255" json:"cover_image,omitempty"`
	IsPublic     bool   `json:"is_public,omitempty"`
	UniversityId uint   `validate:"required,min=1" json:"university_id"`
	CreatorId    uint   `validate:"required,min=1" json:"creator_id"`
}

type UpdateRequest struct {
	Name        string `validate:"required,min=3,max=100" json:"name"`
	Description string `validate:"required,min=10,max=1000" json:"description"`
	Logo        string `validate:"omitempty,url,max=255" json:"logo,omitempty"`
	CoverImage  string `validate:"omitempty,url,max=255" json:"cover_image,omitempty"`
	IsActive    bool   `json:"is_active,omitempty"`
	IsPublic    bool   `json:"is_public,omitempty"`
}

type UpdateStatusRequest struct {
	IsActive bool `json:"is_active"`
	IsPublic bool `json:"is_public"`
}
