package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/service"
	"context"
)

var AdminUser = cAdminUser{}

type cAdminUser struct {
}

// 增加管理员
func (c *cAdminUser) AddAdminUser(ctx context.Context, req *admin.AddAdminUserReq) (res *admin.AdminUserRes, err error) {
	err = service.AdminUser().AddAdminUser(ctx, req)
	return
}

// 更新管理员
func (c *cAdminUser) UpdateAdminUser(ctx context.Context, req *admin.UpdateAdminUserReq) (res *admin.AdminUserRes, err error) {
	err = service.AdminUser().UpdateAdminUser(ctx, req)
	return
}

// 查询管理员
func (c *cAdminUser) GetAdminUser(ctx context.Context, req *admin.GetAdminUserReq) (res *admin.AdminUserRes, err error) {
	res, err = service.AdminUser().GetAdminUser(ctx, req)
	return
}
