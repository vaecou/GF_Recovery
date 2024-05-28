package mini

import "github.com/gogf/gf/v2/frame/g"

type GetBalanceInfoReq struct {
	g.Meta `path:"/balance" tags:"小程序/余额信息" summary:"获取余额信息"`

	Type  int `json:"type" v:"required#类型不能为空"`
	Page  int `json:"page" v:"required#页数不能为空"`
	Limit int `json:"limit" v:"required#页数不能为空"`
}

type BalanceInfoRes struct {
	Balance int               `json:"balance"`
	Freeze  int               `json:"freeze"`
	List    []*BalanceListRes `json:"list"`
}

type BalanceListRes struct {
	Balance   int    `json:"balance"`
	Type      int    `json:"type"`
	CreatedAt string `json:"created_at"`
}

// 提现
type WithdrawReq struct {
	g.Meta `path:"/balance" tags:"小程序/余额信息" summary:"提现"`

	Money int `json:"money" v:"required#提现金额不能为空"`
}
