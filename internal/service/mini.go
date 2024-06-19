// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"GF_Recovery/api/mini"
	"context"
)

type (
	IMiniBalance interface {
		GetUserBalance(ctx context.Context, req *mini.GetBalanceInfoReq) (res []*mini.BalanceListRes, err error)
		// 提现
		Withdraw(ctx context.Context, req *mini.WithdrawReq) (err error)
	}
	IMiniHome interface {
		// 获取使用次数数量
		GetNum(ctx context.Context, req *mini.GetNumReq) (res *mini.GetNumRes, err error)
	}
	IMiniMore interface {
		GetAboutSetting(ctx context.Context, in *mini.GetAboutInfoReq) (res *mini.MoreInfoRes, err error)
		GetProtocoSetting(ctx context.Context, in *mini.GetProtocolInfoReq) (res *mini.MoreInfoRes, err error)
	}
	IMiniOrder interface {
		// 获取type为2的所有balance
		GetUserBalance(ctx context.Context) (res int, err error)
		// 获取type为3的所有kg
		GetUserKG(ctx context.Context) (res float64, err error)
		// 获取订单列表
		GetOrderList(ctx context.Context, req *mini.GetOrderListReq) (res []*mini.OrderListRes, err error)
		// 取消订单
		CancelOrder(ctx context.Context, req *mini.CancelOrderReq) (err error)
		// 删除订单
		DeleteOrder(ctx context.Context, req *mini.DeleteOrderReq) (err error)
		// 创建订单
		AddOrder(ctx context.Context, req *mini.AddOrderReq) (err error)
	}
	IMiniQuestion interface {
		// 随机获取问题列表
		GetQuestion(ctx context.Context, req *mini.GetQuestionReq) (res []*mini.QuestionRes, err error)
	}
	IMiniRegions interface {
		// 获取地址
		GetAddress(ctx context.Context, in *mini.GetAddressReq) (res *mini.AddressRes, err error)
		DeleteAddress(ctx context.Context, in *mini.DeleteAddressReq) (err error)
		UpdateAddress(ctx context.Context, in *mini.UpdateAddressReq) (err error)
		UpdateRadio(ctx context.Context, in *mini.UpdateRadioReq) (err error)
		AddAddress(ctx context.Context, in *mini.AddRegionsReq) (err error)
		GetAddressRadio(ctx context.Context, id int) (res *mini.AddressRes)
		GetAddressList(ctx context.Context, id int) (res []*mini.AddressRes, err error)
	}
	IMiniUsers interface {
		// GetUserBalance
		GetUserBalance(ctx context.Context) (res int, err error)
		// 检查用户status
		CheckUserStatus(ctx context.Context) (res bool, err error)
		// 检查用户是否存在phone
		CheckUserPhone(ctx context.Context) (res bool, err error)
		SaveUserPhone(ctx context.Context, req *mini.SaveUserPhoneReq) (res *mini.UserPhoneRes, err error)
	}
)

var (
	localMiniBalance  IMiniBalance
	localMiniHome     IMiniHome
	localMiniMore     IMiniMore
	localMiniOrder    IMiniOrder
	localMiniQuestion IMiniQuestion
	localMiniRegions  IMiniRegions
	localMiniUsers    IMiniUsers
)

func MiniBalance() IMiniBalance {
	if localMiniBalance == nil {
		panic("implement not found for interface IMiniBalance, forgot register?")
	}
	return localMiniBalance
}

func RegisterMiniBalance(i IMiniBalance) {
	localMiniBalance = i
}

func MiniHome() IMiniHome {
	if localMiniHome == nil {
		panic("implement not found for interface IMiniHome, forgot register?")
	}
	return localMiniHome
}

func RegisterMiniHome(i IMiniHome) {
	localMiniHome = i
}

func MiniMore() IMiniMore {
	if localMiniMore == nil {
		panic("implement not found for interface IMiniMore, forgot register?")
	}
	return localMiniMore
}

func RegisterMiniMore(i IMiniMore) {
	localMiniMore = i
}

func MiniOrder() IMiniOrder {
	if localMiniOrder == nil {
		panic("implement not found for interface IMiniOrder, forgot register?")
	}
	return localMiniOrder
}

func RegisterMiniOrder(i IMiniOrder) {
	localMiniOrder = i
}

func MiniQuestion() IMiniQuestion {
	if localMiniQuestion == nil {
		panic("implement not found for interface IMiniQuestion, forgot register?")
	}
	return localMiniQuestion
}

func RegisterMiniQuestion(i IMiniQuestion) {
	localMiniQuestion = i
}

func MiniRegions() IMiniRegions {
	if localMiniRegions == nil {
		panic("implement not found for interface IMiniRegions, forgot register?")
	}
	return localMiniRegions
}

func RegisterMiniRegions(i IMiniRegions) {
	localMiniRegions = i
}

func MiniUsers() IMiniUsers {
	if localMiniUsers == nil {
		panic("implement not found for interface IMiniUsers, forgot register?")
	}
	return localMiniUsers
}

func RegisterMiniUsers(i IMiniUsers) {
	localMiniUsers = i
}
