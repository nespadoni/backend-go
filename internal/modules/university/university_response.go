package university

import "time"

type Response struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Acronym   string    `json:"acronym,omitempty"`
	City      string    `json:"city,omitempty"`
	State     string    `json:"state,omitempty"`
	Country   string    `json:"country,omitempty"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListResponse struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Acronym string `json:"acronym,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
}
