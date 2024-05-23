package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/service"
	"context"
)

type sMiniQuestion struct {
}

func init() {
	service.RegisterMiniQuestion(NewQuestion())
}

func NewQuestion() *sMiniQuestion {
	return &sMiniQuestion{}
}

// 随机获取问题列表
func (s *sMiniQuestion) GetQuestion(ctx context.Context, req *mini.GetQuestionReq) (res []*mini.GetQuestionRes, err error) {
	err = dao.ReQuestion.Ctx(ctx).OrderRandom().Limit(req.Limit).Scan(&res)
	if err != nil {
		return
	}
	return
}
