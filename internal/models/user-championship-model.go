package models

type Userchampionship struct {
	Base
	ChampionshipID int          `json:"championship_id"`
	Championship   Championship `gorm:"foreignKey:ChampionshipID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MatchID        int          `json:"match_id"`
	Match          Match        `gorm:"foreignKey:MatchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
