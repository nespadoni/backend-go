package lineup

type CreateRequest struct {
	MatchID   uint   `validate:"required" json:"match_id"`
	TeamID    uint   `validate:"required" json:"team_id"`
	PlayerID  uint   `validate:"required" json:"player_id"`
	Position  string `validate:"required,max=50" json:"position"`
	IsStarter bool   `json:"is_starter,omitempty"`
}

type UpdateRequest struct {
	MatchID   uint   `validate:"required" json:"match_id"`
	TeamID    uint   `validate:"required" json:"team_id"`
	PlayerID  uint   `validate:"required" json:"player_id"`
	Position  string `validate:"required,max=50" json:"position"`
	IsStarter bool   `json:"is_starter,omitempty"`
}

type UpdateStatusRequest struct {
	IsStarter bool `json:"is_starter"`
}
