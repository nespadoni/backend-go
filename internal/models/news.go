package models

type News struct {
	Base
	Title       string   `gorm:"size:200;not null" json:"title"`
	Content     string   `gorm:"type:text;not null" json:"content"`
	Summary     string   `gorm:"size:500" json:"summary"` // Resumo para preview
	ImageURL    string   `gorm:"size:255" json:"image_url"`
	IsPublished bool     `gorm:"default:false" json:"is_published"`
	IsPinned    bool     `gorm:"default:false" json:"is_pinned"` // Post fixado
	ViewsCount  int      `gorm:"default:0" json:"views_count"`
	LikesCount  int      `gorm:"default:0" json:"likes_count"`
	AthleticID  int      `json:"athletic_id"`
	Athletic    Athletic `gorm:"foreignKey:AthleticID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AuthorID    int      `json:"author_id"` // Quem postou
	Author      User     `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
