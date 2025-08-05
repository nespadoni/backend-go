package base

import (
	"time"
)

type Base struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"updated_at"`
}
