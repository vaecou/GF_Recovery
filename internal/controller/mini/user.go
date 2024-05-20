package mini

import (
	"GF_Recovery/api/mini"
	"context"
)

var MiniUser = cMiniUser{}

type cMiniUser struct {
}

// 增加用户
func (c *cMiniUser) AddOrUpdateMiniUser(ctx context.Context, req *mini.AddMiniUserReq) (res *mini.MiniUserRes, err error) {
	// sssion, err := service.MiniUser().GetSession(ctx, req)
	// if err != nil {
	// 	return
	// }
	// // 检查是否存在用户
	// if service.MiniUser().CheckAccount(ctx, sssion.UnionID) {

	// } else {

	// }

	return
}
