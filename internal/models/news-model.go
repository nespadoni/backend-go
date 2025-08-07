package models

type News struct {
	Base
	Title      string `gorm:"size:200;not null" json:"title"`
	Content    string `gorm:"type:text;not null" json:"content"`
	AthleticID uint   `json:"athletic_id"` // FK -> Atlética
}
