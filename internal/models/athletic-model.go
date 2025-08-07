package models

type Athletic struct {
	Base
	Name         string     `json:"name"`
	UniversityID int        `json:"university_id"`
	University   University `gorm:"foreignKey:UniversityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Teams        []Team     `gorm:"foreignkey:AthleticID"`
}
