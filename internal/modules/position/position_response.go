package position

type Response struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	SportID uint   `json:"sport_id"`
	Sport   string `json:"sport"`
}
