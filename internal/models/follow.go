package models

type Follow struct {
	Base
	UserID int  `json:"user_id"`
	User   User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// AthleticID uint      `json:"athletic_id"`
	// Athletic   admin.Athletic `gorm:"foreignKey:AthleticID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
