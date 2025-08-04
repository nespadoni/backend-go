package user

import (
	"backend-go/internal/base"
)

type Usuario struct {
	Id       int    `gorm:"primary_key"`
	Nome     string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100;not null"`
	Senha    string `gorm:"size:100;not null"`
	Telefone string `gorm:"size:100;not null"`
	Papel    string `gorm:"size:20;not null"`    // "Administrador" ou "ServicoUsuario"
	AtletaID *uint  `json:"atleta_id,omitempty"` // Se for jogador
	base.Base
}
