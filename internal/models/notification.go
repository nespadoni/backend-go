package models

type Notification struct {
	Base
	UserID  int    `json:"user_id"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title   string `gorm:"size:200;not null" json:"title"`
	Message string `gorm:"type:text;not null" json:"message"`
	Type    string `gorm:"size:50;not null" json:"type"` // "news", "match", "follow", etc.
	IsRead  bool   `gorm:"default:false" json:"is_read"`

	// Referência ao objeto que gerou a notificação
	RelatedID   *int   `json:"related_id"`
	RelatedType string `gorm:"size:50" json:"related_type"` // "athletic", "news", "match"

	AthleticID *int      `json:"athletic_id"` // De qual atlética veio
	Athletic   *Athletic `gorm:"foreignKey:AthleticID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
