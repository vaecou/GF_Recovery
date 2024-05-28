package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sMiniBalance struct {
}

func init() {
	service.RegisterMiniBalance(NewBalance())
}

func NewBalance() *sMiniBalance {
	return &sMiniBalance{}
}

func (c *sMiniBalance) GetUserBalance(ctx context.Context, req *mini.GetBalanceInfoReq) (res []*mini.BalanceListRes, err error) {
	id := gconv.Int(ctx.Value("user_id"))
	if req.Type == 0 {
		err = dao.ReTakeList.Ctx(ctx).Where("user_id", id).Page(req.Page, req.Limit).OrderDesc("created_at").Scan(&res)
		if err != nil {
			return
		}
	} else {
		err = dao.ReTakeList.Ctx(ctx).Where("user_id", id).Where("type", req.Type).Page(req.Page, req.Limit).OrderDesc("created_at").Scan(&res)
		if err != nil {
			return
		}
	}
	return
}

// 提现
func (c *sMiniBalance) Withdraw(ctx context.Context, req *mini.WithdrawReq) (err error) {
	// 开启事务
	dao.ReTakeList.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		// 插入一个提现记录
		id := gconv.Int(ctx.Value("user_id"))
		_, err = dao.ReTakeList.Ctx(ctx).TX(tx).Insert(g.Map{
			"user_id": id,
			"balance": req.Money,
			"type":    2,
		})

		// 减少用户的余额
		_, err = dao.ReUser.Ctx(ctx).TX(tx).Where("id", id).Decrement("balance", req.Money)
		if err != nil {
			return err
		}

		return nil
	})
	return
}
