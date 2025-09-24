package position

type Request struct {
	Name    string `json:"name" binding:"required,min=2,max=50"`
	SportID uint   `json:"sport_id" binding:"required"`
}
