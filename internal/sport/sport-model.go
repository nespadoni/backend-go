package sport

import (
	"backend-go/internal/base"
)

type Sport struct {
	base.Base
	Name string `gorm:"size:50;not null" json:"name"`
}
