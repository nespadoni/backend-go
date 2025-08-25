package championship

import (
	"backend-go/internal/models"
	"time"
)

type Response struct {
	Name       string          `json:"name"`
	StartDate  time.Time       `json:"start-date"`
	EndDate    time.Time       `json:"end-date"`
	AthleticId int             `json:"athletic-id"`
	Athletic   models.Athletic `json:"athletic"`
}
