package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/service"
	"context"
	"fmt"
)

type sMiniMore struct {
}

func init() {
	service.RegisterMiniMore(NewMore())
}

func NewMore() *sMiniMore {
	return &sMiniMore{}
}

func (s *sMiniMore) GetAboutSetting(ctx context.Context, in *mini.GetAboutInfoReq) (res *mini.MoreInfoRes, err error) {
	err = dao.ReSetting.Ctx(ctx).Where("key", "about").Scan(&res)
	if err != nil {
		return
	}
	fmt.Println("res", res)
	return
}
func (s *sMiniMore) GetProtocoSetting(ctx context.Context, in *mini.GetProtocolInfoReq) (res *mini.MoreInfoRes, err error) {
	err = dao.ReSetting.Ctx(ctx).Where("key", "protocol").Scan(&res)
	if err != nil {
		return
	}
	return
}
