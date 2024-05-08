package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

var AdminUser = cAdminUser{}

type cAdminUser struct {
}

// 增加管理员
func (c *cAdminUser) AddAdminUser(ctx context.Context, req *admin.AddAdminUserReq) (res *admin.AdminUserRes, err error) {
	// 检测账号是否已存在
	ok, err := service.AdminUser().CheckAccount(ctx, req.Account)
	if err != nil {
		return
	}
	if ok {
		err = gerror.New("该账号已存在")
		return
	}
	// 增加管理员
	err = service.AdminUser().AddAdminUser(ctx, req)
	if err != nil {
		return
	}
	return
}

// 更新管理员
func (c *cAdminUser) UpdateAdminUser(ctx context.Context, req *admin.UpdateAdminUserReq) (res *admin.AdminUserRes, err error) {
	// 检测账号是否已存在
	ok, err := service.AdminUser().CheckAccount(ctx, req.Account)
	if err != nil {
		return
	}
	if ok {
		err = gerror.New("该账号已存在")
		return
	}
	err = service.AdminUser().UpdateAdminUser(ctx, req)
	if err != nil {
		return
	}
	return
}

// 查询管理员
func (c *cAdminUser) GetAdminUser(ctx context.Context, req *admin.GetAdminUserReq) (res *admin.AdminUserListRes, err error) {
	res, err = service.AdminUser().GetAdminUser(ctx, req)
	if err != nil {
		return
	}
	return
}

// 查询管理员列表
func (c *cAdminUser) GetAdminUserList(ctx context.Context, req *admin.GetAdminUserListReq) (res *admin.AdminUserRes, err error) {
	res = &admin.AdminUserRes{}
	res.Total, err = service.AdminUser().GetAdminUserCount(ctx)
	if err != nil {
		return
	}
	res.List, err = service.AdminUser().GetAdminUserList(ctx, req)
	if err != nil {
		return
	}
	return
}
