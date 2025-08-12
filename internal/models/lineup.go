package models

type Lineup struct {
	Base
	MatchID   uint   `json:"match_id"`
	Match     Match  `gorm:"foreignKey:MatchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TeamID    uint   `json:"team_id"`
	PlayerID  uint   `json:"player_id"`
	Player    Player `gorm:"foreignKey:PlayerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Position  string `gorm:"size:50" json:"position"`
	IsStarter bool   `gorm:"default:true" json:"is_starter"`
}
