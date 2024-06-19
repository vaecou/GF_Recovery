package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/service"
	"context"
)

var More = cMore{}

type cMore struct {
}

func (c *cMore) GetAboutSetting(ctx context.Context, req *mini.GetAboutInfoReq) (res *mini.MoreInfoRes, err error) {
	res, err = service.MiniMore().GetAboutSetting(ctx, req)
	if err != nil {
		return
	}
	return
}

func (c *cMore) GetProtocoSetting(ctx context.Context, req *mini.GetProtocolInfoReq) (res *mini.MoreInfoRes, err error) {
	res, err = service.MiniMore().GetProtocoSetting(ctx, req)
	if err != nil {
		return
	}
	return
}
