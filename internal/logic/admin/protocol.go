package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sProtocol struct {
}

func init() {
	service.RegisterProtocol(NewProtocol())
}

func NewProtocol() *sProtocol {
	return &sProtocol{}
}

// 增加段落
func (s *sProtocol) AddProtocol(ctx context.Context, in *admin.AddProtocolReq) (err error) {
	_, err = dao.ReProtocol.Ctx(ctx).Insert(in)
	return
}

// 更新段落
func (s *sProtocol) UpdateProtocol(ctx context.Context, in *admin.UpdateProtocolReq) (err error) {
	affected, err := dao.ReProtocol.Ctx(ctx).OmitEmpty().Where("id", in.ID).UpdateAndGetAffected(in)
	if err != nil {
		return
	}
	if affected == 0 {
		err = gerror.New("该段落不存在")
	}
	return
}

// 获取回收协议
func (s *sProtocol) GetProtocol(ctx context.Context, in *admin.GetProtocolReq) (res *admin.ProtocolListRes, err error) {
	err = dao.ReProtocol.Ctx(ctx).Where("id", in.ID).Scan(&res)
	if err != nil {
		return
	}
	return
}

// 查询回收协议列表数量
func (s *sProtocol) GetProtocolCount(ctx context.Context) (res int, err error) {
	res, err = dao.ReProtocol.Ctx(ctx).Count()
	if err != nil {
		return
	}
	return
}

// 获取回收协议列表
func (s *sProtocol) GetProtocolList(ctx context.Context, in *admin.GetProtocolListReq) (res []*admin.ProtocolListRes, err error) {
	err = dao.ReProtocol.Ctx(ctx).Page(in.Page, in.Limit).Scan(&res)
	if err != nil {
		return
	}
	return
}

// 删除回收协议
func (s *sProtocol) DeleteProtocol(ctx context.Context, in *admin.DeleteProtocolReq) (err error) {
	_, err = dao.ReProtocol.Ctx(ctx).Delete("id", in.ID)
	return
}
