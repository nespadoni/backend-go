package models

type User struct {
	Id        int    `gorm:"primary_key"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;not null"`
	Password  string `gorm:"size:100;not null"`
	Phone     string `gorm:"size:100;not null"`
	Role      string `gorm:"size:20;not null"`     // "Admin" ou "user-service"
	AthleteID *uint  `json:"athlete_id,omitempty"` // Se for jogador
	Base
}
