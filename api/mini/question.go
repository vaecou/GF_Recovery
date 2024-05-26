package mini

import "github.com/gogf/gf/v2/frame/g"

type GetQuestionReq struct {
	g.Meta `path:"/question" tags:"小程序/问题" summary:"随机获取问题列表"`

	Limit int `json:"limit"`
}
type QuestionRes struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created"`
}
