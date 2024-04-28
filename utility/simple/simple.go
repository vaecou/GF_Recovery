package simple

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/util/gconv"
)

// SafeGo 安全的调用协程，遇到错误时输出错误日志而不是抛出panic
func SafeGo(ctx context.Context, f func(ctx context.Context), lv ...interface{}) {
	var level = glog.LEVEL_ERRO
	if len(lv) > 0 {
		level = gconv.Int(lv[0])
	}

	err := grpool.AddWithRecover(ctx, func(ctx context.Context) {
		f(ctx)
	}, func(ctx context.Context, err error) {
		Logf(level, ctx, "SafeGo安全执行失败：%+v", err)
	})

	if err != nil {
		Logf(level, ctx, "SafeGo添加并恢复错误：%+v", err)
		return
	}
}

// Logf 日志输出
func Logf(level int, ctx context.Context, format string, v ...interface{}) {
	switch level {
	case glog.LEVEL_DEBU:
		g.Log().Debugf(ctx, format, v...)
	case glog.LEVEL_INFO:
		g.Log().Infof(ctx, format, v...)
	case glog.LEVEL_NOTI:
		g.Log().Noticef(ctx, format, v...)
	case glog.LEVEL_WARN:
		g.Log().Warningf(ctx, format, v...)
	case glog.LEVEL_ERRO:
		g.Log().Errorf(ctx, format, v...)
	case glog.LEVEL_CRIT:
		g.Log().Criticalf(ctx, format, v...)
	case glog.LEVEL_PANI:
		g.Log().Panicf(ctx, format, v...)
	case glog.LEVEL_FATA:
		g.Log().Fatalf(ctx, format, v...)
	default:
		g.Log().Errorf(ctx, format, v...)
	}
}
