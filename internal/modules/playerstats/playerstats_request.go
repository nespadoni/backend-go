package playerstats

type CreateRequest struct {
	PlayerID      uint `validate:"required" json:"player_id"`
	MatchID       uint `validate:"required" json:"match_id"`
	Goals         int  `validate:"min=0" json:"goals"`
	Assists       int  `validate:"min=0" json:"assists"`
	YellowCards   int  `validate:"min=0,max=2" json:"yellow_cards"`
	RedCards      int  `validate:"min=0,max=1" json:"red_cards"`
	MinutesPlayed int  `validate:"min=0,max=120" json:"minutes_played"`
}

type UpdateRequest struct {
	PlayerID      uint `validate:"required" json:"player_id"`
	MatchID       uint `validate:"required" json:"match_id"`
	Goals         int  `validate:"min=0" json:"goals"`
	Assists       int  `validate:"min=0" json:"assists"`
	YellowCards   int  `validate:"min=0,max=2" json:"yellow_cards"`
	RedCards      int  `validate:"min=0,max=1" json:"red_cards"`
	MinutesPlayed int  `validate:"min=0,max=120" json:"minutes_played"`
}
