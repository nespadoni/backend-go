package models

type Team struct {
	Base
	Name       string   `gorm:"size:100;not null" json:"name"`
	AthleticID uint     `json:"athletic_id"` // FK -> Atlética
	Athletic   Athletic `gorm:"foreignKey:AthleticID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SportID    *uint    `json:"sport_id,omitempty"` // FK -> Sport (opcional)
	Sport      *Sport   `gorm:"foreignKey:SportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"sport,omitempty"`
}
