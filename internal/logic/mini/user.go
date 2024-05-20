package mini

import (
	"GF_Recovery/internal/service"
)

type sMiniUser struct {
}

func init() {
	service.RegisterMiniUser(NewUser())
}

func NewUser() *sMiniUser {
	return &sMiniUser{}
}
