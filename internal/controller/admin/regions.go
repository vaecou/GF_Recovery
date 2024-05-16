package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/service"
	"context"
)

var Regions = cRegions{}

type cRegions struct {
}

// 增加地区
func (c *cRegions) AddRegions(ctx context.Context, req *admin.AddRegionsReq) (res *admin.RegionsRes, err error) {
	err = service.Regions().AddRegions(ctx, req)
	if err != nil {
		return
	}
	return
}

// 更新地区
func (c *cRegions) UpdateRegions(ctx context.Context, req *admin.UpdateRegionsReq) (res *admin.RegionsRes, err error) {
	err = service.Regions().UpdateRegions(ctx, req)
	if err != nil {
		return
	}
	return
}

// 查询地区
func (c *cRegions) GetRegions(ctx context.Context, req *admin.GetRegionsReq) (res *admin.RegionsRes, err error) {
	res, err = service.Regions().GetRegions(ctx, req)
	if err != nil {
		return
	}
	return
}

// 查询地区列表
func (c *cRegions) GetRegionsList(ctx context.Context, req *admin.GetRegionsListReq) (res []*admin.RegionsRes, err error) {
	res, err = service.Regions().GetRegionsList(ctx, req)
	if err != nil {
		return
	}
	return
}

// 删除地区
func (c *cRegions) DeleteRegions(ctx context.Context, req *admin.DeleteRegionsReq) (res *admin.RegionsRes, err error) {
	err = service.Regions().DeleteRegions(ctx, req)
	if err != nil {
		return
	}
	return
}
