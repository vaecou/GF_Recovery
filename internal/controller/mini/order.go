package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/service"
	"context"
)

var Order = cOrder{}

type cOrder struct {
}

// 获取订单列表
func (c *cOrder) GetOrderList(ctx context.Context, req *mini.GetOrderListReq) (res []*mini.OrderListRes, err error) {
	res, err = service.MiniOrder().GetOrderList(ctx, req)
	if err != nil {
		return
	}
	return
}

// 取消订单
func (c *cOrder) CancelOrder(ctx context.Context, req *mini.CancelOrderReq) (res *mini.OrderListRes, err error) {
	err = service.MiniOrder().CancelOrder(ctx, req)
	if err != nil {
		return
	}
	return
}

// 删除订单
func (c *cOrder) DeleteOrder(ctx context.Context, req *mini.DeleteOrderReq) (res *mini.OrderListRes, err error) {
	err = service.MiniOrder().DeleteOrder(ctx, req)
	if err != nil {
		return
	}
	return
}

// 创建订单
func (c *cOrder) AddOrder(ctx context.Context, req *mini.AddOrderReq) (res *mini.OrderListRes, err error) {
	err = service.MiniOrder().AddOrder(ctx, req)
	if err != nil {
		return
	}
	return
}
