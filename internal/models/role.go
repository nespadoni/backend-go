package models

type Role struct {
	Base
	Name        string `gorm:"size:100;not null" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
	Admin       bool   `json:"admin"`
}
