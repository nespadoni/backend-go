package models

import "time"

type Sport struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:50;not null" json:"name"`
	IconURL   string    `json:"icon_url"`
	Teams     []Team    `gorm:"foreignKey:SportID" json:"teams"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
