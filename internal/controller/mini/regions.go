package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

var Regions = cRegions{}

type cRegions struct {
}

func (c *cRegions) GetAddress(ctx context.Context, req *mini.GetAddressReq) (res *mini.AddressRes, err error) {
	res, err = service.MiniRegions().GetAddress(ctx, req)
	if err != nil {
		return
	}
	return
}

func (c *cRegions) DeleteAddress(ctx context.Context, req *mini.DeleteAddressReq) (res *mini.AddressRes, err error) {
	err = service.MiniRegions().DeleteAddress(ctx, req)
	if err != nil {
		return
	}
	return
}

func (c *cRegions) UpdateAddress(ctx context.Context, req *mini.UpdateAddressReq) (res *mini.AddressRes, err error) {
	err = service.MiniRegions().UpdateAddress(ctx, req)
	if err != nil {
		return
	}
	return
}

func (c *cRegions) UpdateRadio(ctx context.Context, req *mini.UpdateRadioReq) (res *mini.AddressRes, err error) {
	err = service.MiniRegions().UpdateRadio(ctx, req)
	if err != nil {
		return
	}
	return
}

// 查询地区
func (c *cRegions) AddAddress(ctx context.Context, req *mini.AddRegionsReq) (res []*mini.QuestionRes, err error) {
	err = service.MiniRegions().AddAddress(ctx, req)
	if err != nil {
		return
	}
	return
}

// 查询地址
func (c *cRegions) GetAddressList(ctx context.Context, req *mini.GetAddressListReq) (res *mini.AddressListRes, err error) {
	res = &mini.AddressListRes{}
	id := gconv.Int(ctx.Value("user_id"))
	res.Radio, err = service.MiniRegions().GetAddressRadio(ctx, id)
	if err != nil {
		res.Radio = 0
	}
	res.List, err = service.MiniRegions().GetAddressList(ctx, id)
	if err != nil {
		return
	}
	return
}
