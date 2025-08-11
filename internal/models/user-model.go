package models

type User struct {
	Base
	Name         string     `gorm:"size:100;not null" json:"name"`
	Email        string     `gorm:"size:100;not null" json:"email"`
	Password     string     `gorm:"size:100;not null" json:"password"`
	Telephone    string     `gorm:"size:100;not null" json:"telephone"`
	UniversityID int        `json:"university_id"`
	University   University `gorm:"foreignKey:UniversityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID       int        `json:"role_id"`
	Role         Role       `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
