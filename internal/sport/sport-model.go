package sport

import (
	"backend-go/internal/base"
	"backend-go/internal/team"
)

type Sport struct {
	ID      uint        `gorm:"primaryKey" json:"id"`
	Name    string      `gorm:"size:50;not null" json:"name"`
	IconURL string      `json:"icon_url"`
	Teams   []team.Team `gorm:"foreignKey:SportID" json:"teams"`
	base.Base
}
