package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sAbout struct {
}

func init() {
	service.RegisterAbout(NewAbout())
}

func NewAbout() *sAbout {
	return &sAbout{}
}

// 增加段落
func (s *sAbout) AddAbout(ctx context.Context, in *admin.AddAboutReq) (err error) {
	_, err = dao.ReAbout.Ctx(ctx).Insert(in)
	return
}

// 更新段落
func (s *sAbout) UpdateAbout(ctx context.Context, in *admin.UpdateAboutReq) (err error) {
	affected, err := dao.ReAbout.Ctx(ctx).OmitEmpty().Where("id", in.ID).UpdateAndGetAffected(in)
	if err != nil {
		return
	}
	if affected == 0 {
		err = gerror.New("该段落不存在")
	}
	return
}

// 获取关于我们
func (s *sAbout) GetAbout(ctx context.Context, in *admin.GetAboutReq) (res *admin.AboutListRes, err error) {
	err = dao.ReAbout.Ctx(ctx).Where("id", in.ID).Scan(&res)
	if err != nil {
		return
	}
	return
}

// 查询关于我们列表数量
func (s *sAbout) GetAboutCount(ctx context.Context) (res int, err error) {
	res, err = dao.ReAbout.Ctx(ctx).Count()
	if err != nil {
		return
	}
	return
}

// 获取关于我们列表
func (s *sAbout) GetAboutList(ctx context.Context, in *admin.GetAboutListReq) (res []*admin.AboutListRes, err error) {
	err = dao.ReAbout.Ctx(ctx).Page(in.Page, in.Limit).Scan(&res)
	if err != nil {
		return
	}
	return
}

// 删除关于我们
func (s *sAbout) DeleteAbout(ctx context.Context, in *admin.DeleteAboutReq) (err error) {
	_, err = dao.ReAbout.Ctx(ctx).Delete("id", in.ID)
	return
}
