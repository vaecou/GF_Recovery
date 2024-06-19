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
	IAbout interface {
		// 增加段落
		AddAbout(ctx context.Context, in *admin.AddAboutReq) (err error)
		// 更新段落
		UpdateAbout(ctx context.Context, in *admin.UpdateAboutReq) (err error)
		// 获取关于我们
		GetAbout(ctx context.Context, in *admin.GetAboutReq) (res *admin.AboutListRes, err error)
		// 查询关于我们列表数量
		GetAboutCount(ctx context.Context) (res int, err error)
		// 获取关于我们列表
		GetAboutList(ctx context.Context, in *admin.GetAboutListReq) (res []*admin.AboutListRes, err error)
		// 删除关于我们
		DeleteAbout(ctx context.Context, in *admin.DeleteAboutReq) (err error)
	}
	IMiniUser interface {
		// 查询用户列表数量
		GetMiniUserCount(ctx context.Context, in *admin.GetMiniUserListReq) (int, error)
		// 查询用户列表
		GetMiniUserList(ctx context.Context, in *admin.GetMiniUserListReq) ([]*admin.MiniUserListRes, error)
		// 更新用户
		UpdateMiniUser(ctx context.Context, in *admin.UpdateMiniUserReq) (err error)
	}
	IProtocol interface {
		// 增加段落
		AddProtocol(ctx context.Context, in *admin.AddProtocolReq) (err error)
		// 更新段落
		UpdateProtocol(ctx context.Context, in *admin.UpdateProtocolReq) (err error)
		// 获取回收协议
		GetProtocol(ctx context.Context, in *admin.GetProtocolReq) (res *admin.ProtocolListRes, err error)
		// 查询回收协议列表数量
		GetProtocolCount(ctx context.Context) (res int, err error)
		// 获取回收协议列表
		GetProtocolList(ctx context.Context, in *admin.GetProtocolListReq) (res []*admin.ProtocolListRes, err error)
		// 删除回收协议
		DeleteProtocol(ctx context.Context, in *admin.DeleteProtocolReq) (err error)
	}
	IQuestion interface {
		// 查询问题是否存在
		CheckQuestion(ctx context.Context, title string) (ok bool)
		// 增加问题
		AddQuestion(ctx context.Context, in *admin.AddQuestionReq) (err error)
		// 更新问题
		UpdateQuestion(ctx context.Context, in *admin.UpdateQuestionReq) (err error)
		// 获取问题
		GetQuestion(ctx context.Context, in *admin.GetQuestionReq) (res *admin.QuestionListRes, err error)
		// 查询问题列表数量
		GetQuestionCount(ctx context.Context) (res int, err error)
		// 获取问题列表
		GetQuestionList(ctx context.Context, in *admin.GetQuestionListReq) (res []*admin.QuestionListRes, err error)
		// 删除问题
		DeleteQuestion(ctx context.Context, in *admin.DeleteQuestionReq) (err error)
	}
	IRegions interface {
		// 增加地区
		AddRegions(ctx context.Context, in *admin.AddRegionsReq) (err error)
		// 更新地区
		UpdateRegions(ctx context.Context, in *admin.UpdateRegionsReq) (err error)
		// 删除地区
		DeleteRegions(ctx context.Context, in *admin.DeleteRegionsReq) (err error)
		// 获取地区
		GetRegions(ctx context.Context, in *admin.GetRegionsReq) (res *admin.RegionsRes, err error)
		// 获取地区列表
		GetRegionsList(ctx context.Context, in *admin.GetRegionsListReq) (res []*admin.RegionsRes, err error)
	}
	ISetting interface {
		// 检查是否存在Key
		CheckKey(ctx context.Context, key string) (ok bool)
		// 添加设置到数据库
		AddSetting(ctx context.Context, in *admin.AddOrUpdateSettingReq) (err error)
		// 更新设置到数据库
		UpdateSetting(ctx context.Context, in *admin.AddOrUpdateSettingReq) (err error)
		// 查询设置
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
	localAbout     IAbout
	localMiniUser  IMiniUser
	localProtocol  IProtocol
	localQuestion  IQuestion
	localRegions   IRegions
	localSetting   ISetting
	localAdminUser IAdminUser
)

func About() IAbout {
	if localAbout == nil {
		panic("implement not found for interface IAbout, forgot register?")
	}
	return localAbout
}

func RegisterAbout(i IAbout) {
	localAbout = i
}

func MiniUser() IMiniUser {
	if localMiniUser == nil {
		panic("implement not found for interface IMiniUser, forgot register?")
	}
	return localMiniUser
}

func RegisterMiniUser(i IMiniUser) {
	localMiniUser = i
}

func Protocol() IProtocol {
	if localProtocol == nil {
		panic("implement not found for interface IProtocol, forgot register?")
	}
	return localProtocol
}

func RegisterProtocol(i IProtocol) {
	localProtocol = i
}

func Question() IQuestion {
	if localQuestion == nil {
		panic("implement not found for interface IQuestion, forgot register?")
	}
	return localQuestion
}

func RegisterQuestion(i IQuestion) {
	localQuestion = i
}

func Regions() IRegions {
	if localRegions == nil {
		panic("implement not found for interface IRegions, forgot register?")
	}
	return localRegions
}

func RegisterRegions(i IRegions) {
	localRegions = i
}

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
