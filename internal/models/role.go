package models

type Role struct {
	Base
	Name        string    `gorm:"size:100;not null" json:"name"`
	Description string    `gorm:"size:255;not null" json:"description"`
	Level       int       `gorm:"default:0" json:"level"`         // Hierarquia de permissões (0-100)
	Permissions string    `gorm:"type:json" json:"permissions"`   // JSON com permissões específicas
	IsSystem    bool      `gorm:"default:false" json:"is_system"` // Roles do sistema vs custom
	AthleticID  *int      `json:"athletic_id"`                    // Para roles específicas de uma atlética
	Athletic    *Athletic `gorm:"foreignKey:AthleticID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
