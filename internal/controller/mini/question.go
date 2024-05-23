package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/service"
	"context"
)

var Question = cQuestion{}

type cQuestion struct {
}

// 随机获取问题列表
func (c *cQuestion) GetQuestion(ctx context.Context, req *mini.GetQuestionReq) (res []*mini.GetQuestionRes, err error) {
	// 如果没填limit就默认为3
	if req.Limit == 0 {
		req.Limit = 3
	}
	res, err = service.MiniQuestion().GetQuestion(ctx, req)
	if err != nil {
		return
	}
	return
}
