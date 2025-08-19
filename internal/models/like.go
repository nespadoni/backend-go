package models

type Like struct {
	Base
	UserID int  `json:"user_id"`
	User   User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Usando polimorfismo para diferentes tipos de conteúdo
	LikeableID   int    `json:"likeable_id"`
	LikeableType string `gorm:"size:50" json:"likeable_type"` // "news", "comment", etc.
}
