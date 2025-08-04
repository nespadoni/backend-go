package player

import (
	"backend-go/internal/base"
)

type Player struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"size:100;not null" json:"name"`
	TeamID   uint   `json:"team_id"`
	Position string `gorm:"size:50" json:"position"` // Ex: "Atacante", "Goleiro"
	UserID   *uint  `json:"user_id,omitempty"`       // Se for vinculado a um usuário
	base.Base
}
