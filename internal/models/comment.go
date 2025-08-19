package models

type Comment struct {
	Base
	Content string `gorm:"type:text;not null" json:"content"`
	UserID  int    `json:"user_id"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Para diferentes tipos de conteúdo
	CommentableID   int    `json:"commentable_id"`
	CommentableType string `gorm:"size:50" json:"commentable_type"` // "news", "match", etc.

	// Para threads de comentários
	ParentID *int     `json:"parent_id"`
	Parent   *Comment `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	LikesCount int `gorm:"default:0" json:"likes_count"`
}
