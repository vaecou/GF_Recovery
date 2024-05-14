package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sQuestion struct {
}

func init() {
	service.RegisterQuestion(NewQuestion())
}

func NewQuestion() *sQuestion {
	return &sQuestion{}
}

// 查询问题是否存在
func (s *sQuestion) CheckQuestion(ctx context.Context, title string) (ok bool) {
	count, err := dao.ReQuestion.Ctx(ctx).Where(dao.ReQuestion.Columns().Title, title).Count()
	if err != nil {
		return false
	}
	if count > 0 {
		return true
	}
	return
}

// 增加问题
func (s *sQuestion) AddQuestion(ctx context.Context, in *admin.AddQuestionReq) (err error) {
	_, err = dao.ReQuestion.Ctx(ctx).Insert(in)
	return
}

// 更新问题
func (s *sQuestion) UpdateQuestion(ctx context.Context, in *admin.UpdateQuestionReq) (err error) {
	affected, err := dao.ReQuestion.Ctx(ctx).OmitEmpty().Where("id", in.ID).UpdateAndGetAffected(in)
	if err != nil {
		return
	}
	if affected == 0 {
		err = gerror.New("该问题不存在")
	}
	return
}

// 查询管理员列表数量
func (s *sQuestion) GetQuestionCount(ctx context.Context) (res int, err error) {
	res, err = dao.ReQuestion.Ctx(ctx).Count()
	if err != nil {
		return
	}
	return
}

// 获取问题列表
func (s *sQuestion) GetQuestionList(ctx context.Context, in *admin.GetQuestionListReq) (res []*admin.QuestionListRes, err error) {
	err = dao.ReQuestion.Ctx(ctx).Page(in.Page, in.Limit).Scan(&res)
	if err != nil {
		return
	}
	return
}

// 删除问题
func (s *sQuestion) DeleteQuestion(ctx context.Context, in *admin.DeleteQuestionReq) (err error) {
	_, err = dao.ReQuestion.Ctx(ctx).Delete("id", in.ID)
	return
}
