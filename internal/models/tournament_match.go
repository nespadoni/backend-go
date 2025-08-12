package models

type TournamentMatch struct {
	Base
	TournamentId int        `json:"tournament_id"`
	Tournament   Tournament `gorm:"foreignKey:TournamentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MatchID      int        `json:"match_id"`
	Match        Match      `gorm:"foreignKey:MatchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
