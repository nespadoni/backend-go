package follow

import (
	"backend-go/internal/admin"
	"backend-go/internal/base"
	"backend-go/internal/user"
)

type Follow struct {
	base.Base
	UserID     uint           `json:"user_id"`
	User       user.Usuario   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AthleticID uint           `json:"athletic_id"`
	Athletic   admin.Athletic `gorm:"foreignKey:AthleticID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
