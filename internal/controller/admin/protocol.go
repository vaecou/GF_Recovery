package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/service"
	"context"
)

var Protocol = cProtocol{}

type cProtocol struct {
}

// 增加回收协议
func (c *cProtocol) AddProtocol(ctx context.Context, req *admin.AddProtocolReq) (res *admin.ProtocolListRes, err error) {
	// 增加段落
	err = service.Protocol().AddProtocol(ctx, req)
	if err != nil {
		return
	}
	return
}

// 更新回收协议
func (c *cProtocol) UpdateProtocol(ctx context.Context, req *admin.UpdateProtocolReq) (res *admin.ProtocolListRes, err error) {
	err = service.Protocol().UpdateProtocol(ctx, req)
	if err != nil {
		return
	}
	return
}

// 查询回收协议
func (c *cProtocol) GetProtocol(ctx context.Context, req *admin.GetProtocolReq) (res *admin.ProtocolListRes, err error) {
	res, err = service.Protocol().GetProtocol(ctx, req)
	if err != nil {
		return
	}
	return
}

// 获取回收协议列表
func (c *cProtocol) GetProtocolList(ctx context.Context, req *admin.GetProtocolListReq) (res *admin.ProtocolRes, err error) {
	res = &admin.ProtocolRes{}
	res.Total, err = service.Protocol().GetProtocolCount(ctx)
	if err != nil {
		return
	}
	res.List, err = service.Protocol().GetProtocolList(ctx, req)
	if err != nil {
		return
	}
	return
}

// 删除回收协议
func (c *cProtocol) DeleteProtocol(ctx context.Context, req *admin.DeleteProtocolReq) (res *admin.ProtocolListRes, err error) {
	err = service.Protocol().DeleteProtocol(ctx, req)
	if err != nil {
		return
	}
	return
}
