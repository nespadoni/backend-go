package university

type Request struct {
	Name     string `validate:"required,min=3,max=200" json:"name"`
	Acronym  string `validate:"max=20" json:"acronym,omitempty"`
	City     string `validate:"max=100" json:"city,omitempty"`
	State    string `validate:"max=50" json:"state,omitempty"`
	Country  string `validate:"max=50" json:"country,omitempty"`
	IsActive bool   `json:"is_active,omitempty"`
}
