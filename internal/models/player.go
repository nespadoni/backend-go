package models

type Player struct {
	Base
	Name       string   `gorm:"size:100;not null" json:"name"`
	TeamID     int      `json:"team_id"`
	Team       Team     `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PositionID int      `json:"position_id"`
	Position   Position `gorm:"foreignKey:PositionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID     *int     `json:"user_id"`
	User       *User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
