package models

type Notification struct {
	Base
	UserID  int    `json:"user_id"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // ADICIONAR
	Title   string `gorm:"size:200;not null" json:"title"`
	Message string `gorm:"type:text;not null" json:"message"`
	Type    string `gorm:"size:50;not null" json:"type"`
	IsRead  bool   `gorm:"default:false" json:"is_read"`
}
