package models

type Follow struct {
	Base
	UserID     int      `json:"user_id"`
	User       User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AthleticID int      `json:"athletic_id"`
	Athletic   Athletic `gorm:"foreignKey:AthleticID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
