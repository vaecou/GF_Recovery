package mini

import "github.com/gogf/gf/v2/frame/g"

type GetAddressReq struct {
	g.Meta `path:"/address" tags:"小程序/地址" summary:"获取地址"`

	ID int `json:"id"`
}

type DeleteAddressReq struct {
	g.Meta `path:"/address" tags:"小程序/地址" summary:"删除地址"`

	ID int `json:"id"`
}

type UpdateAddressReq struct {
	g.Meta `path:"/address" tags:"小程序/地址" summary:"更新地址"`

	ID        int    `json:"id"`
	Name      string `json:"name" v:"required#联系人不能为空"`
	Phone     string `json:"phone" v:"required#手机号不能为空"`
	Provinces string `json:"provinces" v:"required#省不能为空"`
	Citys     string `json:"citys" v:"required#市不能为空"`
	Areas     string `json:"areas" v:"required#区不能为空"`
	Detail    string `json:"detail" v:"required#详细地址不能为空"`
	Default   bool   `json:"default"`
}

type UpdateRadioReq struct {
	g.Meta `path:"/address/radio" tags:"小程序/地址" summary:"更新默认地址"`

	ID int `json:"id"`
}

type GetAddressRadioReq struct {
	g.Meta `path:"/address/radio" tags:"小程序/地址" summary:"获取默认地址"`
}

type AddRegionsReq struct {
	g.Meta `path:"/address" tags:"小程序/地址" summary:"新增地址"`

	UserID    int
	Name      string `json:"name" v:"required#联系人不能为空"`
	Phone     string `json:"phone" v:"required#手机号不能为空"`
	Provinces string `json:"provinces" v:"required#省不能为空"`
	Citys     string `json:"citys" v:"required#市不能为空"`
	Areas     string `json:"areas" v:"required#区不能为空"`
	Detail    string `json:"detail" v:"required#详细地址不能为空"`
	Default   bool   `json:"default"`
}

type GetAddressListReq struct {
	g.Meta `path:"/address/list" tags:"小程序/地址" summary:"获取地址列表"`
}

type AddressListRes struct {
	Radio *AddressRes   `json:"radio"`
	List  []*AddressRes `json:"list"`
}

type AddressRes struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Provinces string `json:"provinces"`
	Citys     string `json:"citys"`
	Areas     string `json:"areas"`
	Detail    string `json:"detail"`
	Default   bool   `json:"default"`
}

type GetRegionsListReq struct {
	g.Meta `path:"/regions/list" tags:"小程序/地区" summary:"获取地区列表"`
}

type RegionsRes struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Sort     int    `json:"sort"`
	Regions  string `json:"regions"`

	Children []*RegionsRes `json:"children"`
}
