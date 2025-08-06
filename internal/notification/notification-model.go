package notification

import (
	"backend-go/internal/base"
	"time"
)

type Notification struct {
	base.Base
	UserID    uint      `json:"user_id"`
	Title     string    `gorm:"size:200;not null" json:"title"`
	Message   string    `gorm:"type:text;not null" json:"message"`
	Type      string    `gorm:"size:50;not null" json:"type"` // "match_result", "news", "schedule_change"
	IsRead    bool      `gorm:"default:false" json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

// SERA QUE FUNCIONOU?
