package models

type Sport struct {
	Base
	Name      string     `gorm:"size:50;not null" json:"name"`
	Positions []Position `gorm:"foreignkey:SportID"`
}
