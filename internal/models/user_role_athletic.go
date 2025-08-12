package models

import "time"

//Tabela associativa para relacionar a permissão do usuário na atlética
//Para fazer igual a uma pagina do facebook

type UserRoleAthletic struct {
	UserID     int        `gorm:"primaryKey;uniqueIndex:user_role_athletic_idx" json:"user_id"`
	User       User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID     int        `gorm:"primaryKey;uniqueIndex:user_role_athletic_idx" json:"role_id"`
	Role       Role       `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AthleticID *int       `gorm:"uniqueIndex:user_role_athletic_idx" json:"athletic_id"`
	Athletic   *Athletic  `gorm:"foreignKey:AthleticID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
