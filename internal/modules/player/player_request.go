package player

type CreateRequest struct {
	Name       string `validate:"required,min=2,max=100" json:"name"`
	TeamID     int    `validate:"required" json:"team_id"`
	PositionID int    `validate:"required" json:"position_id"`
	UserID     *int   `json:"user_id,omitempty"`
}

type UpdateRequest struct {
	Name       string `validate:"required,min=2,max=100" json:"name"`
	TeamID     int    `validate:"required" json:"team_id"`
	PositionID int    `validate:"required" json:"position_id"`
	UserID     *int   `json:"user_id,omitempty"`
}

type UpdateTeamRequest struct {
	TeamID int `validate:"required" json:"team_id"`
}

type UpdatePositionRequest struct {
	PositionID int `validate:"required" json:"position_id"`
}
