package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/service"
	"GF_Recovery/utility/generate"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sMiniOrder struct {
}

func init() {
	service.RegisterMiniOrder(NewOrder())
}

func NewOrder() *sMiniOrder {
	return &sMiniOrder{}
}

// 获取type为2的所有balance
func (c *sMiniOrder) GetUserBalance(ctx context.Context) (res int, err error) {
	id := gconv.Int(ctx.Value("user_id"))
	balance, err := dao.ReOrderList.Ctx(ctx).Where("user_id", id).Where("type", 2).Sum("balance")
	res = gconv.Int(balance)
	return
}

// 获取type为3的所有kg
func (c *sMiniOrder) GetUserKG(ctx context.Context) (res float64, err error) {
	id := gconv.Int(ctx.Value("user_id"))
	res, err = dao.ReOrderList.Ctx(ctx).Where("user_id", id).Where("type", 3).Sum("kg")
	return
}

// 获取订单列表
func (c *sMiniOrder) GetOrderList(ctx context.Context, req *mini.GetOrderListReq) (res []*mini.OrderListRes, err error) {
	id := gconv.Int(ctx.Value("user_id"))
	if req.Type == 0 {
		err = dao.ReOrderList.Ctx(ctx).Where("user_id", id).Page(req.Page, req.Limit).OrderDesc("updated_at").Scan(&res)
		if err != nil {
			return
		}
	} else {
		err = dao.ReOrderList.Ctx(ctx).Where("user_id", id).Where("type", req.Type).Page(req.Page, req.Limit).OrderDesc("updated_at").Scan(&res)
		if err != nil {
			return
		}
	}
	return
}

// 取消订单
func (c *sMiniOrder) CancelOrder(ctx context.Context, req *mini.CancelOrderReq) (err error) {
	_, err = dao.ReOrderList.Ctx(ctx).Where("order_id", req.OrderID).Where("type", 1).Update(g.Map{
		"type": 0,
	})
	if err != nil {
		return
	}
	return
}

// 删除订单
func (c *sMiniOrder) DeleteOrder(ctx context.Context, req *mini.DeleteOrderReq) (err error) {
	_, err = dao.ReOrderList.Ctx(ctx).Where("order_id", req.OrderID).Where("type", 0).Delete()
	if err != nil {
		return
	}
	return
}

// 创建订单
func (c *sMiniOrder) AddOrder(ctx context.Context, req *mini.AddOrderReq) (err error) {
	id := gconv.Int(ctx.Value("user_id"))
	req.OrderID, err = generate.GenerateOrderNo("1")
	if err != nil {
		return
	}
	req.UserID = id
	_, err = dao.ReOrderList.Ctx(ctx).Insert(req)
	if err != nil {
		return
	}
	return
}
