// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ReAddressList is the golang structure for table re_address_list.
type ReAddressList struct {
	UserId    int    `json:"user_id"   orm:"user_id"   ` // 用户ID
	Id        int    `json:"id"        orm:"id"        ` // ID
	Name      string `json:"name"      orm:"name"      ` // 联系人
	Phone     string `json:"phone"     orm:"phone"     ` // 手机号
	Provinces string `json:"provinces" orm:"provinces" ` // 省
	Citys     string `json:"citys"     orm:"citys"     ` // 市
	Areas     string `json:"areas"     orm:"areas"     ` // 区
	Detail    string `json:"detail"    orm:"detail"    ` // 详细地址
	Default   bool   `json:"default"   orm:"default"   ` // 是否默认
}
