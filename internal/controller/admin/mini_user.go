package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/service"
	"context"
)

var MiniUser = cMiniUser{}

type cMiniUser struct {
}

// 查询管理员列表
func (c *cMiniUser) GetMiniUserList(ctx context.Context, req *admin.GetMiniUserListReq) (res *admin.MiniUserRes, err error) {
	res = &admin.MiniUserRes{}
	res.Total, err = service.MiniUser().GetMiniUserCount(ctx)
	if err != nil {
		return
	}
	res.List, err = service.MiniUser().GetMiniUserList(ctx, req)
	if err != nil {
		return
	}
	return
}
