package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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
	r, err := g.Client().Post(
		ctx,
		"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=6ff0b65c-9ee6-472a-af87-f23ec7c2064f",
		`{ 
			"msgtype": "text", 
			"text": { 
				"content": "此订单已被客户取消，请注意留意！", 
				"mentioned_list":["@all"] 
			},
		}`,
	)
	if err != nil {
		return
	}
	defer r.Close()
	r, err = g.Client().Post(
		ctx,
		"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=6ff0b65c-9ee6-472a-af87-f23ec7c2064f",
		`{ 
			"msgtype": "markdown", 
			"markdown": { 
				"content": "# 订单编号：<font color=\"warning\">`+gconv.String(req.OrderID)+`</font>" 
			},
		}`,
	)
	if err != nil {
		return
	}
	defer r.Close()
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
	ok, err := service.MiniUsers().CheckUserStatus(ctx)
	if err != nil {
		return
	}
	if ok {
		err = gerror.New("您已被系统锁定")
		return
	}
	ok, err = service.MiniUsers().CheckUserPhone(ctx)
	if err != nil {
		return
	}
	if !ok {
		err = gerror.New("请先绑定手机号")
		return
	}
	err = service.MiniOrder().AddOrder(ctx, req)
	if err != nil {
		return
	}

	r, err := g.Client().Post(
		ctx,
		"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=6ff0b65c-9ee6-472a-af87-f23ec7c2064f",
		`{ 
			"msgtype": "text", 
			"text": { 
				"content": "来新订单了，请注意查收！", 
				"mentioned_list":["@all"] 
			},
		}`,
	)
	if err != nil {
		return
	}
	defer r.Close()
	r, err = g.Client().Post(
		ctx,
		"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=6ff0b65c-9ee6-472a-af87-f23ec7c2064f",
		`{ 
			"msgtype": "markdown", 
			"markdown": { 
				"content": "# 订单编号：<font color=\"warning\">`+gconv.String(req.OrderID)+`</font>\n
				> 联系姓名：<font color=\"comment\">`+gconv.String(req.Name)+`</font>\n
				> 联系电话：<font color=\"comment\">`+gconv.String(req.Phone)+`</font>\n
				> 预估重量：<font color=\"comment\">`+gconv.String(req.Estimate)+`</font>\n
				> 预约时间：<font color=\"comment\">`+gconv.String(req.Day)+`（`+gconv.String(req.Week)+`）`+gconv.String(req.Time)+`</font>\n
				> 取件地址：<font color=\"comment\">`+gconv.String(req.Provinces)+gconv.String(req.Citys)+gconv.String(req.Areas)+gconv.String(req.Detail)+`</font>\n
				> 打开订单：<font color=\"comment\">[请点击我](http://work.weixin.qq.com/api/doc)</font>" 
			},
		}`,
	)
	if err != nil {
		return
	}
	defer r.Close()

	return
}
