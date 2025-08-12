package models

type Player struct {
	Base
	Name     string `gorm:"size:100;not null" json:"name"`
	TeamID   uint   `json:"team_id"`
	Team     Team   `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Position string `gorm:"size:50" json:"position"`
	UserID   *int   `json:"user_id"`
	User     *User  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
