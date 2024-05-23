package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/service"
	"context"
)

type sMiniHome struct {
}

func init() {
	service.RegisterMiniHome(NewHome())
}

func NewHome() *sMiniHome {
	return &sMiniHome{}
}

// 获取使用次数数量
func (s *sMiniHome) GetNum(ctx context.Context, req *mini.GetNumReq) (res *mini.GetNumRes, err error) {
	res = &mini.GetNumRes{}
	res.Num, err = dao.ReOrderList.Ctx(ctx).Count()
	if err != nil {
		return
	}
	return
}
