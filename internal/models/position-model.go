package models

type Position struct {
	Base
	Name    string `json:"name"`
	SportID int    `json:"sport_id"`
	Sport   Sport  `gorm:"foreignKey:SportID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
