// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ReRegions is the golang structure for table re_regions.
type ReRegions struct {
	Id       int    `json:"id"        orm:"id"        ` // ID
	ParentId int    `json:"parent_id" orm:"parent_id" ` // 父ID
	Sort     int    `json:"sort"      orm:"sort"      ` // 排序
	Regions  string `json:"regions"   orm:"regions"   ` // 值
}
