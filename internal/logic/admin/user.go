package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/grand"
)

type sAdminUser struct {
}

func init() {
	service.RegisterAdminUser(NewUser())
}

func NewUser() *sAdminUser {
	return &sAdminUser{}
}

// 检查账号是否重复
func (s *sAdminUser) CheckAccount(ctx context.Context, account string) (bool, error) {
	count, err := dao.ReAdminUser.Ctx(ctx).Where(dao.ReAdminUser.Columns().Account, account).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// 增加管理员
func (s *sAdminUser) AddAdminUser(ctx context.Context, in *admin.AddAdminUserReq) (err error) {
	in.Salt = grand.S(8)
	in.Password = gmd5.MustEncryptString(in.Password + in.Salt)
	_, err = dao.ReAdminUser.Ctx(ctx).Insert(in)
	return
}

// 更新管理员
func (s *sAdminUser) UpdateAdminUser(ctx context.Context, in *admin.UpdateAdminUserReq) (err error) {
	fmt.Println("in", in)
	if in.Password != "" {
		in.Salt = grand.S(8)
		in.Password = gmd5.MustEncryptString(in.Password + in.Salt)
	}
	affected, err := dao.ReAdminUser.Ctx(ctx).OmitEmpty().Where("id", in.ID).UpdateAndGetAffected(in)
	if err != nil {
		return
	}
	if affected == 0 {
		err = gerror.New("该管理员不存在")
	}
	return
}

// 查询管理员
func (s *sAdminUser) GetAdminUser(ctx context.Context, in *admin.GetAdminUserReq) (res *admin.AdminUserListRes, err error) {
	err = dao.ReAdminUser.Ctx(ctx).Where("id", in.ID).Scan(&res)
	if err != nil {
		return
	}
	return
}

// 查询管理员列表数量
func (s *sAdminUser) GetAdminUserCount(ctx context.Context) (res int, err error) {
	res, err = dao.ReAdminUser.Ctx(ctx).Count()
	if err != nil {
		return
	}
	return
}

// 查询管理员列表
func (s *sAdminUser) GetAdminUserList(ctx context.Context, in *admin.GetAdminUserListReq) (res []*admin.AdminUserListRes, err error) {
	err = dao.ReAdminUser.Ctx(ctx).Page(in.Page, in.Limit).Scan(&res)
	if err != nil {
		return
	}
	return
}
