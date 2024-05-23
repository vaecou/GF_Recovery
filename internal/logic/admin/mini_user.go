package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/service"
	"context"
)

type sMiniUser struct {
}

func init() {
	service.RegisterMiniUser(NewMiniUser())
}

func NewMiniUser() *sMiniUser {
	return &sMiniUser{}
}

// 查询管理员列表数量
func (s *sMiniUser) GetMiniUserCount(ctx context.Context) (res int, err error) {
	res, err = dao.ReUser.Ctx(ctx).Count()
	if err != nil {
		return
	}
	return
}

// 查询管理员列表
func (s *sMiniUser) GetMiniUserList(ctx context.Context, in *admin.GetMiniUserListReq) (res []*admin.MiniUserListRes, err error) {
	err = dao.ReUser.Ctx(ctx).Page(in.Page, in.Limit).Scan(&res)
	if err != nil {
		return
	}
	return
}
