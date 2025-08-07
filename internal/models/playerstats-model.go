package models

type PlayerStats struct {
	Base
	PlayerID      uint   `json:"player_id"`
	Player        Player `gorm:"foreignKey:PlayerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MatchID       uint   `json:"match_id"`
	Match         Match  `gorm:"foreignKey:MatchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Goals         int    `gorm:"default:0" json:"goals"`
	Assists       int    `gorm:"default:0" json:"assists"`
	YellowCards   int    `gorm:"default:0" json:"yellow_cards"`
	RedCards      int    `gorm:"default:0" json:"red_cards"`
	MinutesPlayed int    `gorm:"default:0" json:"minutes_played"`
}
