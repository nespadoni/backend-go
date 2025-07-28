package models

import (
	"time"
)

type User struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Email     string    `gorm:"size:100;not null" json:"email"`
	Password  string    `gorm:"size:100;not null" json:"-"`
	Phone     string    `gorm:"size:100;not null" json:"phone"`
	Role      string    `gorm:"size:20;not null" json:"role"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
