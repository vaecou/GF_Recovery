package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/model/do"
	"GF_Recovery/internal/service"
	"context"
)

type sSetting struct {
}

func init() {
	service.RegisterSetting(NewSetting())
}

func NewSetting() *sSetting {
	return &sSetting{}
}

// 检查是否存在Key
func (s *sSetting) CheckKey(ctx context.Context, key string) (ok bool) {
	var res int
	res, err := dao.ReSetting.Ctx(ctx).Where(dao.ReSetting.Columns().Key, key).Count()
	if err != nil {
		return false
	}
	return res > 0
}

// 添加设置到数据库
func (s *sSetting) AddSetting(ctx context.Context, in *admin.AddOrUpdateSettingReq) (err error) {
	setting := &do.ReSetting{
		Key:   in.Key,
		Value: in.Value,
	}
	_, err = dao.ReSetting.Ctx(ctx).Insert(setting)
	return
}

// 更新设置到数据库
func (s *sSetting) UpdateSetting(ctx context.Context, in *admin.AddOrUpdateSettingReq) (err error) {
	_, err = dao.ReSetting.Ctx(ctx).Where("key", in.Key).Update(do.ReSetting{
		Value: in.Value,
	})
	return
}

// 查询设置
func (s *sSetting) GetSetting(ctx context.Context, in *admin.GetSettingReq) (res *admin.SettingRes, err error) {
	err = dao.ReSetting.Ctx(ctx).Where("key", in.Key).Scan(&res)
	if err != nil {
		return
	}
	return
}
