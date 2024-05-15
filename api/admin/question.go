package admin

import "github.com/gogf/gf/v2/frame/g"

type AddQuestionReq struct {
	g.Meta `path:"/question" tags:"问题" summary:"增加问题"`

	Title   string `json:"title" v:"required#标题不能为空"`
	Content string `json:"content" v:"required#内容不能为空"`
}

type UpdateQuestionReq struct {
	g.Meta `path:"/question" tags:"问题" summary:"更新问题"`

	ID      int    `json:"id" v:"required#ID不能为空"`
	Title   string `json:"title" v:"required#标题不能为空"`
	Content string `json:"content" v:"required#内容不能为空"`
}

type GetQuestionReq struct {
	g.Meta `path:"/question" tags:"问题" summary:"查询问题"`

	ID int `json:"id" v:"required#ID不能为空"`
}

type GetQuestionListReq struct {
	g.Meta `path:"/question/list" tags:"问题" summary:"查询问题列表"`

	Page  int `json:"page" v:"required#页不能为空"`
	Limit int `json:"limit" v:"required#数量不能为空"`
}

type DeleteQuestionReq struct {
	g.Meta `path:"/question" tags:"问题" summary:"删除问题"`

	ID int `json:"id" v:"required#ID不能为空"`
}

type QuestionRes struct {
	Total int                `json:"total"`
	List  []*QuestionListRes `json:"list"`
}

type QuestionListRes struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created"`
}
