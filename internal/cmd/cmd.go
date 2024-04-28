package cmd

import (
	"context"
	"time"

	"GF_Recovery/internal/consts"
	"GF_Recovery/utility/simple"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gsession"
)

var (
	Main = &gcmd.Command{
		Description: "默认启动所有服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			return All.Func(ctx, parser)
		},
	}

	All = gcmd.Command{
		Description: "启动所有服务器的命令条目",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Debug(ctx, "正在启动所有服务")

			// 需要启动的服务
			allServers := []*gcmd.Command{Http}
			for _, server := range allServers {
				var cmd = server
				simple.SafeGo(ctx, func(ctx context.Context) {
					if err := cmd.Func(ctx, parser); err != nil {
						g.Log().Fatalf(ctx, "%v启动失败:%v", cmd.Name, err)
					}
				})
			}

			// 信号监听
			signalListen(ctx, signalHandlerForOverall)

			<-serverCloseSignal
			serverWg.Wait()
			g.Log().Debug(ctx, "所有服务已成功关闭")
			return
		},
	}

	Http = &gcmd.Command{
		Name:        "Http",
		Description: "网络交互相关服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化http服务
			s := g.Server()

			// Session配置
			s.SetSessionMaxAge(time.Hour)                                                  // 1小时的Session有效期
			s.SetSessionStorage(gsession.NewStorageRedis(g.Redis(), consts.SessionPrefix)) // 使用Redis存储Session

			s.Run()

			return
		},
	}
)
