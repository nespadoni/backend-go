package sport

type CreateRequest struct {
	Name         string `validate:"required,min=2,max=100" json:"name"`
	Description  string `validate:"max=500" json:"description,omitempty"`
	Abbreviation string `validate:"max=10" json:"abbreviation,omitempty"`
	Icon         string `validate:"omitempty,url,max=255" json:"icon,omitempty"`
	Category     string `validate:"max=50" json:"category,omitempty"`
	MinPlayers   int    `validate:"min=1,max=50" json:"min_players"`
	MaxPlayers   int    `validate:"min=1,max=50" json:"max_players"`
	Rules        string `validate:"max=2000" json:"rules,omitempty"`
	IsPopular    bool   `json:"is_popular,omitempty"`
}

type UpdateRequest struct {
	Name         string `validate:"required,min=2,max=100" json:"name"`
	Description  string `validate:"max=500" json:"description,omitempty"`
	Abbreviation string `validate:"max=10" json:"abbreviation,omitempty"`
	Icon         string `validate:"omitempty,url,max=255" json:"icon,omitempty"`
	Category     string `validate:"max=50" json:"category,omitempty"`
	MinPlayers   int    `validate:"min=1,max=50" json:"min_players"`
	MaxPlayers   int    `validate:"min=1,max=50" json:"max_players"`
	Rules        string `validate:"max=2000" json:"rules,omitempty"`
	IsActive     bool   `json:"is_active,omitempty"`
	IsPopular    bool   `json:"is_popular,omitempty"`
}

type UpdateStatusRequest struct {
	IsActive  bool `json:"is_active"`
	IsPopular bool `json:"is_popular"`
}
