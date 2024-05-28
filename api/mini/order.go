package mini

import "github.com/gogf/gf/v2/frame/g"

// 获取订单列表
type GetOrderListReq struct {
	g.Meta `path:"/order/list" tags:"小程序/订单" summary:"获取订单列表"`

	Type  int `json:"type" v:"required#类型不能为空"`
	Page  int `json:"page" v:"required#页数不能为空"`
	Limit int `json:"limit" v:"required#页数不能为空"`
}

// 取消订单
type CancelOrderReq struct {
	g.Meta `path:"/order" tags:"小程序/订单" summary:"取消订单"`

	OrderID int `json:"order_id" v:"required#订单ID不能为空"`
}

// 删除订单
type DeleteOrderReq struct {
	g.Meta `path:"/order" tags:"小程序/订单" summary:"删除订单"`

	OrderID int `json:"order_id" v:"required#订单ID不能为空"`
}

// 创建订单
type AddOrderReq struct {
	g.Meta `path:"/order" tags:"小程序/订单" summary:"增加订单"`

	OrderID   int64  `json:"order_id"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name" v:"required#联系人不能为空"`
	Phone     string `json:"phone" v:"required#手机号不能为空"`
	Provinces string `json:"provinces" v:"required#省不能为空"`
	Citys     string `json:"citys" v:"required#市不能为空"`
	Areas     string `json:"areas" v:"required#区不能为空"`
	Detail    string `json:"detail" v:"required#详细地址不能为空"`
	Day       string `json:"day" v:"required#日期不能为空"`
	Week      string `json:"week" v:"required#星期不能为空"`
	Time      string `json:"time" v:"required#时间不能为空"`
	Type      int    `json:"type" v:"required#类型不能为空"`
	Estimate  string `json:"estimate" v:"required#预估重量不能为空"`
}

// 订单列表
type OrderListRes struct {
	OrderID   int64   `json:"order_id"`
	UserID    int     `json:"user_id"`
	RecylerID int     `json:"recycler_id"`
	KG        float64 `json:"kg"`
	Balance   int     `json:"balance"`
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	Provinces string  `json:"provinces"`
	Citys     string  `json:"citys"`
	Areas     string  `json:"areas"`
	Detail    string  `json:"detail"`
	Day       string  `json:"day"`
	Week      string  `json:"week"`
	Time      string  `json:"time"`
	Type      int     `json:"type"`
	Estimate  string  `json:"estimate"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string  `json:"deleted_at"`
}
