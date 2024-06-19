package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/service"
	"context"
)

var About = cAbout{}

type cAbout struct {
}

// 增加关于我们
func (c *cAbout) AddAbout(ctx context.Context, req *admin.AddAboutReq) (res *admin.AboutListRes, err error) {
	// 增加段落
	err = service.About().AddAbout(ctx, req)
	if err != nil {
		return
	}
	return
}

// 更新关于我们
func (c *cAbout) UpdateAbout(ctx context.Context, req *admin.UpdateAboutReq) (res *admin.AboutListRes, err error) {
	err = service.About().UpdateAbout(ctx, req)
	if err != nil {
		return
	}
	return
}

// 查询关于我们
func (c *cAbout) GetAbout(ctx context.Context, req *admin.GetAboutReq) (res *admin.AboutListRes, err error) {
	res, err = service.About().GetAbout(ctx, req)
	if err != nil {
		return
	}
	return
}

// 获取关于我们列表
func (c *cAbout) GetAboutList(ctx context.Context, req *admin.GetAboutListReq) (res *admin.AboutRes, err error) {
	res = &admin.AboutRes{}
	res.Total, err = service.About().GetAboutCount(ctx)
	if err != nil {
		return
	}
	res.List, err = service.About().GetAboutList(ctx, req)
	if err != nil {
		return
	}
	return
}

// 删除关于我们
func (c *cAbout) DeleteAbout(ctx context.Context, req *admin.DeleteAboutReq) (res *admin.AboutListRes, err error) {
	err = service.About().DeleteAbout(ctx, req)
	if err != nil {
		return
	}
	return
}
