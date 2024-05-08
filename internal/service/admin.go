// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"GF_Recovery/api/admin"
	"context"
)

type (
	IAdminUser interface {
		// CheckAccount
		CheckAccount(ctx context.Context, account string) (bool, error)
		// 增加管理员
		AddAdminUser(ctx context.Context, in *admin.AddAdminUserReq) (err error)
		// 更新管理员
		UpdateAdminUser(ctx context.Context, in *admin.UpdateAdminUserReq) (err error)
		// 查询管理员
		GetAdminUser(ctx context.Context, in *admin.GetAdminUserReq) (res *admin.AdminUserListRes, err error)
		// 查询管理员列表数量
		GetAdminUserCount(ctx context.Context) (res int, err error)
		// 查询管理员列表
		GetAdminUserList(ctx context.Context, in *admin.GetAdminUserListReq) (res []*admin.AdminUserListRes, err error)
	}
)

var (
	localAdminUser IAdminUser
)

func AdminUser() IAdminUser {
	if localAdminUser == nil {
		panic("implement not found for interface IAdminUser, forgot register?")
	}
	return localAdminUser
}

func RegisterAdminUser(i IAdminUser) {
	localAdminUser = i
}
