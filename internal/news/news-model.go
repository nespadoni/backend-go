package news

import (
	"backend-go/internal/base"
)

type News struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Title      string `gorm:"size:200;not null" json:"title"`
	Content    string `gorm:"type:text;not null" json:"content"`
	AthleticID uint   `json:"athletic_id"` // FK -> Atlética
	base.Base
}
