package models

type University struct {
	Base
	Name     string `gorm:"size:200;not null;unique" json:"name"`
	Acronym  string `gorm:"size:20" json:"acronym,omitempty"`
	City     string `gorm:"size:100" json:"city,omitempty"`
	State    string `gorm:"size:50" json:"state,omitempty"`
	Country  string `gorm:"size:50;default:'Brasil'" json:"country,omitempty"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
}
