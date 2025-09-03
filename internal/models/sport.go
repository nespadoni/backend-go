package models

type Sport struct {
	Base
	Name         string `gorm:"size:100;not null;uniqueIndex" json:"name"`
	Description  string `gorm:"type:text" json:"description,omitempty"`
	Abbreviation string `gorm:"size:10" json:"abbreviation,omitempty"` // Ex: "FUT" para Futebol
	Icon         string `gorm:"size:255" json:"icon,omitempty"`        // URL do ícone do esporte
	Category     string `gorm:"size:50" json:"category,omitempty"`     // Ex: "Coletivo", "Individual"
	MinPlayers   int    `gorm:"default:1" json:"min_players"`          // Mínimo de jogadores
	MaxPlayers   int    `gorm:"default:1" json:"max_players"`          // Máximo de jogadores
	IsActive     bool   `gorm:"default:true" json:"is_active"`
	IsPopular    bool   `gorm:"default:false" json:"is_popular"`  // Destacar esportes populares
	Rules        string `gorm:"type:text" json:"rules,omitempty"` // Regras específicas

	// Relacionamentos
	Positions   []Position   `gorm:"foreignKey:SportID" json:"positions,omitempty"`
	Tournaments []Tournament `gorm:"foreignKey:SportID" json:"tournaments,omitempty"`
	Teams       []Team       `gorm:"foreignKey:SportID" json:"teams,omitempty"`
	Matches     []Match      `gorm:"foreignKey:SportID" json:"matches,omitempty"`
}
