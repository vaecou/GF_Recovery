package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/service"
	"context"
)

var Home = cHome{}

type cHome struct {
}

// 获取使用次数数量
func (c *cHome) GetNum(ctx context.Context, req *mini.GetNumReq) (res *mini.GetNumRes, err error) {
	res, err = service.MiniHome().GetNum(ctx, req)
	if err != nil {
		return
	}
	return
}
