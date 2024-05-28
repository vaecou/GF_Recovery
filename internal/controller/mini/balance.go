package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var Balance = cBalance{}

type cBalance struct {
}

func (c *cBalance) GetBalanceInfo(ctx context.Context, req *mini.GetBalanceInfoReq) (res *mini.BalanceInfoRes, err error) {
	res = &mini.BalanceInfoRes{}
	res.Balance, err = service.MiniUsers().GetUserBalance(ctx)
	if err != nil {
		return
	}
	res.Freeze, err = service.MiniOrder().GetUserBalance(ctx)
	if err != nil {
		return
	}
	res.List, err = service.MiniBalance().GetUserBalance(ctx, req)
	if err != nil {
		return
	}
	return
}

// 提现
func (c *cBalance) Withdraw(ctx context.Context, req *mini.WithdrawReq) (res *mini.BalanceInfoRes, err error) {
	balance, err := service.MiniUsers().GetUserBalance(ctx)
	if err != nil {
		return
	}
	if balance < req.Money {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "余额不足")
		return
	}

	err = service.MiniBalance().Withdraw(ctx, req)
	if err != nil {
		return
	}

	return
}
