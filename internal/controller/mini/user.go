package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/service"
	"context"
)

var MiniUsers = cMiniUsers{}

type cMiniUsers struct {
}

func (c *cMiniUsers) CheckUserPhone(ctx context.Context, req *mini.GetUserPhoneReq) (res *mini.UserPhoneRes, err error) {
	res = &mini.UserPhoneRes{}
	res.Result, err = service.MiniUsers().CheckUserPhone(ctx)
	if err != nil {
		return
	}
	return
}

// SaveUserPhone
func (c *cMiniUsers) SaveUserPhone(ctx context.Context, req *mini.SaveUserPhoneReq) (res *mini.UserPhoneRes, err error) {
	res, err = service.MiniUsers().SaveUserPhone(ctx, req)
	if err != nil {
		return
	}
	return
}
