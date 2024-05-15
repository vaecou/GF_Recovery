package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

var Question = cQuestion{}

type cQuestion struct {
}

// 增加问题
func (c *cQuestion) AddQuestion(ctx context.Context, req *admin.AddQuestionReq) (res *admin.QuestionListRes, err error) {
	ok := service.Question().CheckQuestion(ctx, req.Title)
	if ok {
		err = gerror.New("该问题已存在")
		return
	}
	// 增加管理员
	err = service.Question().AddQuestion(ctx, req)
	if err != nil {
		return
	}
	return
}

// 更新问题
func (c *cQuestion) UpdateQuestion(ctx context.Context, req *admin.UpdateQuestionReq) (res *admin.QuestionListRes, err error) {
	err = service.Question().UpdateQuestion(ctx, req)
	if err != nil {
		return
	}
	return
}

// 查询问题
func (c *cQuestion) GetQuestion(ctx context.Context, req *admin.GetQuestionReq) (res *admin.QuestionListRes, err error) {
	res, err = service.Question().GetQuestion(ctx, req)
	if err != nil {
		return
	}
	return
}

// 获取问题列表
func (c *cQuestion) GetQuestionList(ctx context.Context, req *admin.GetQuestionListReq) (res *admin.QuestionRes, err error) {
	res = &admin.QuestionRes{}
	res.Total, err = service.Question().GetQuestionCount(ctx)
	if err != nil {
		return
	}
	res.List, err = service.Question().GetQuestionList(ctx, req)
	if err != nil {
		return
	}
	return
}

// 删除问题
func (c *cQuestion) DeleteQuestion(ctx context.Context, req *admin.DeleteQuestionReq) (res *admin.QuestionListRes, err error) {
	err = service.Question().DeleteQuestion(ctx, req)
	if err != nil {
		return
	}
	return
}
