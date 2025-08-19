package models

type Athletic struct {
	Base
	Name           string     `gorm:"size:100;not null" json:"name"`
	Description    string     `gorm:"type:text" json:"description"`
	Logo           string     `gorm:"size:255" json:"logo"`        // URL do logo
	CoverImage     string     `gorm:"size:255" json:"cover_image"` // Imagem de capa
	IsActive       bool       `gorm:"default:true" json:"is_active"`
	IsPublic       bool       `gorm:"default:true" json:"is_public"`    // Pública ou privada
	FollowersCount int        `gorm:"default:0" json:"followers_count"` // Cache de seguidores
	UniversityID   int        `json:"university_id"`
	University     University `gorm:"foreignKey:UniversityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatorID      int        `json:"creator_id"` // Quem criou a atlética
	Creator        User       `gorm:"foreignKey:CreatorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
