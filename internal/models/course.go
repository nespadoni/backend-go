package models

type Course struct {
	Base
	Name         string     `gorm:"size:150;not null" json:"name"`
	Code         string     `gorm:"size:20" json:"code,omitempty"`             // Ex: "CC", "ENG", "MED"
	Duration     int        `gorm:"default:4" json:"duration"`                 // Anos de duração
	Level        string     `gorm:"size:20;default:'graduation'" json:"level"` // 'graduation', 'postgraduate', 'master', 'phd'
	UniversityID int        `gorm:"not null" json:"university_id"`
	University   University `gorm:"foreignKey:UniversityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"university"`
	IsActive     bool       `gorm:"default:true" json:"is_active"`
}
