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
	ISetting interface {
		// 检查是否存在Key
		CheckKey(ctx context.Context, key string) (ok bool)
		// 添加关于我们到数据库
		AddSetting(ctx context.Context, in *admin.AddOrUpdateSettingReq) (err error)
		// 更新关于我们到数据库
		UpdateSetting(ctx context.Context, in *admin.AddOrUpdateSettingReq) (err error)
		// 查询关于我们
		GetSetting(ctx context.Context, in *admin.GetSettingReq) (res *admin.SettingRes, err error)
	}
	IAdminUser interface {
		// 查询账号是否存在
		CheckAccount(ctx context.Context, account string) (ok bool)
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
	localSetting   ISetting
	localAdminUser IAdminUser
)

func Setting() ISetting {
	if localSetting == nil {
		panic("implement not found for interface ISetting, forgot register?")
	}
	return localSetting
}

func RegisterSetting(i ISetting) {
	localSetting = i
}

func AdminUser() IAdminUser {
	if localAdminUser == nil {
		panic("implement not found for interface IAdminUser, forgot register?")
	}
	return localAdminUser
}

func RegisterAdminUser(i IAdminUser) {
	localAdminUser = i
}
