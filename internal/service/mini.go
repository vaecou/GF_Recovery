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
		GetQuestion(ctx context.Context, req *mini.GetQuestionReq) (res []*mini.GetQuestionRes, err error)
	}
)

var (
	localMiniHome     IMiniHome
	localMiniQuestion IMiniQuestion
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
