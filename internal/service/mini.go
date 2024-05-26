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
	IMiniHome interface {
		// 获取使用次数数量
		GetNum(ctx context.Context, req *mini.GetNumReq) (res *mini.GetNumRes, err error)
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
)

var (
	localMiniHome     IMiniHome
	localMiniQuestion IMiniQuestion
	localMiniRegions  IMiniRegions
)

func MiniHome() IMiniHome {
	if localMiniHome == nil {
		panic("implement not found for interface IMiniHome, forgot register?")
	}
	return localMiniHome
}

func RegisterMiniHome(i IMiniHome) {
	localMiniHome = i
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
