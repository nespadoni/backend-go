package models

type Athletic struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"size:100;not null" json:"name"`
	University string `gorm:"size:150;not null" json:"university"`
	LogoURL    string `json:"logo_url"`
	AdminID    uint   `json:"admin_id"` // FK -> Usuário admin
	Teams      []Team `gorm:"foreignKey:AthleticID" json:"teams"`
	Base
}
