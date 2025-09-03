package models

type Position struct {
	Base
	Name    string `gorm:"size:50;not null" json:"name"`
	SportID uint   `json:"sport_id"` // FK -> Sport
	Sport   Sport  `gorm:"foreignKey:SportID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
