package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sMiniUser struct {
}

func init() {
	service.RegisterMiniUser(NewMiniUser())
}

func NewMiniUser() *sMiniUser {
	return &sMiniUser{}
}

// 查询用户列表数量
func (s *sMiniUser) GetMiniUserCount(ctx context.Context, in *admin.GetMiniUserListReq) (int, error) {
	// 初始化查询对象
	query := dao.ReUser.Ctx(ctx)

	// 根据筛选条件构建查询
	if in.Title != "" {
		switch in.Select {
		case "":
			// 筛选全部: 从ID和手机号中筛选in.Title
			query = query.Where("id LIKE ? OR phone LIKE ?", "%"+in.Title+"%", "%"+in.Title+"%")
		case "1":
			// 筛选ID
			query = query.Where("id LIKE ?", "%"+in.Title+"%")
		case "2":
			// 筛选手机号
			query = query.Where("phone LIKE ?", "%"+in.Title+"%")
		}
	}

	// 执行计数查询并返回结果
	count, err := query.Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 查询用户列表
func (s *sMiniUser) GetMiniUserList(ctx context.Context, in *admin.GetMiniUserListReq) ([]*admin.MiniUserListRes, error) {
	var res []*admin.MiniUserListRes

	// 初始化查询对象
	query := dao.ReUser.Ctx(ctx)

	// 根据筛选条件构建查询
	if in.Title != "" {
		switch in.Select {
		case "":
			// 筛选全部: 从ID和手机号中筛选in.Title
			query = query.Where("id LIKE ? OR phone LIKE ?", "%"+in.Title+"%", "%"+in.Title+"%")
		case "1":
			// 筛选ID
			query = query.Where("id LIKE ?", "%"+in.Title+"%")
		case "2":
			// 筛选手机号
			query = query.Where("phone LIKE ?", "%"+in.Title+"%")
		}
	}

	// 分页查询并扫描结果
	err := query.Page(in.Page, in.Limit).Scan(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 更新用户
func (s *sMiniUser) UpdateMiniUser(ctx context.Context, in *admin.UpdateMiniUserReq) (err error) {
	affected, err := dao.ReUser.Ctx(ctx).Where("id", in.ID).UpdateAndGetAffected(in)
	if err != nil {
		return
	}
	if affected == 0 {
		err = gerror.New("该用户不存在")
	}
	return
}
