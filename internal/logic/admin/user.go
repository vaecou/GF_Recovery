package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/service"
	"context"

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

// 增加管理员
func (s *sAdminUser) AddAdminUser(ctx context.Context, in *admin.AddAdminUserReq) (err error) {
	in.Salt = grand.S(8)
	in.Password = gmd5.MustEncryptString(in.Password + in.Salt)
	_, err = dao.ReAdminUser.Ctx(ctx).Insert(in)
	return
}

// 更新管理员
func (s *sAdminUser) UpdateAdminUser(ctx context.Context, in *admin.UpdateAdminUserReq) (err error) {
	in.Salt = grand.S(8)
	in.Password = gmd5.MustEncryptString(in.Password + in.Salt)
	affected, err := dao.ReAdminUser.Ctx(ctx).Where("id", in.Id).UpdateAndGetAffected(in)
	if err != nil {
		return
	}
	if affected == 0 {
		err = gerror.New("该管理员不存在")
	}
	return
}

// 查询管理员
func (s *sAdminUser) GetAdminUser(ctx context.Context, in *admin.GetAdminUserReq) (res *admin.AdminUserRes, err error) {
	err = dao.ReAdminUser.Ctx(ctx).Where("id", in.Id).Scan(&res)
	if err != nil {
		return
	}
	// 检查用户是否存在
	if res == nil {
		err = gerror.New("该管理员不存在")
	}
	return
}
