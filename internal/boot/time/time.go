package time

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 设置时区
func SetTimeZone(ctx context.Context) {
	// 默认上海时区
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		g.Log().Fatalf(ctx, "SetTimeZone时区设置异常：%+v", err)
		return
	}
}
