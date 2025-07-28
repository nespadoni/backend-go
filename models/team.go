package models

type Team struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	Name       string   `gorm:"size:100;not null" json:"name"`
	AthleticID uint     `json:"athletic_id"` // FK -> Atlética
	SportID    uint     `json:"sport_id"`    // FK -> Esporte
	Players    []Player `gorm:"foreignKey:TeamID" json:"players"`
	Base
}
