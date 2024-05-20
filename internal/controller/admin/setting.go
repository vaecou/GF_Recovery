package admin

import (
	"GF_Recovery/api/admin"
	"GF_Recovery/internal/service"
	"context"
	"fmt"
)

var Setting = cSetting{}

type cSetting struct {
}

// 增加关于我们
func (c *cSetting) AddOrUpdateSetting(ctx context.Context, req *admin.AddOrUpdateSettingReq) (res *admin.SettingRes, err error) {
	// 检查是否存在key
	if service.Setting().CheckKey(ctx, req.Key) {
		// 更新
		err = service.Setting().UpdateSetting(ctx, req)
		if err != nil {
			return nil, err
		}
	} else {
		err = service.Setting().AddSetting(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return
}

// 查询关于我们
func (c *cSetting) GetSetting(ctx context.Context, req *admin.GetSettingReq) (res *admin.SettingRes, err error) {
	res, err = service.Setting().GetSetting(ctx, req)
	fmt.Println("res", res)
	if err != nil {
		return
	}
	return
}
