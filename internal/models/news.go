package models

type News struct {
	Base
	Title      string   `gorm:"size:200;not null" json:"title"`
	Content    string   `gorm:"type:text;not null" json:"content"`
	AthleticID int      `json:"athletic_id"`
	Athletic   Athletic `gorm:"foreignKey:AthleticID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
