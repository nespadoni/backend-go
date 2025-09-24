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

	// Informações adicionais
	Nickname string  `gorm:"size:50" json:"nickname"`
	Height   float32 `json:"height"`
	Weight   float32 `json:"weight"`
	Status   string  `gorm:"size:100" json:"status"`
	Notes    string  `gorm:"size:500" json:"notes"`
}
