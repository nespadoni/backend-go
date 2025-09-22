package result

type CreateRequest struct {
	MatchID    int    `validate:"required" json:"match_id"`
	TeamAScore int    `validate:"min=0" json:"team_a_score"`
	TeamBScore int    `validate:"min=0" json:"team_b_score"`
	Status     string `validate:"required,oneof=scheduled live finished postponed" json:"status"`
	IsLive     bool   `json:"is_live,omitempty"`
}

type UpdateRequest struct {
	MatchID    int    `validate:"required" json:"match_id"`
	TeamAScore int    `validate:"min=0" json:"team_a_score"`
	TeamBScore int    `validate:"min=0" json:"team_b_score"`
	Status     string `validate:"required,oneof=scheduled live finished postponed" json:"status"`
	IsLive     bool   `json:"is_live,omitempty"`
}

type UpdateScoreRequest struct {
	TeamAScore int `validate:"min=0" json:"team_a_score"`
	TeamBScore int `validate:"min=0" json:"team_b_score"`
}

type UpdateStatusRequest struct {
	Status string `validate:"required,oneof=scheduled live finished postponed" json:"status"`
	IsLive bool   `json:"is_live"`
}
