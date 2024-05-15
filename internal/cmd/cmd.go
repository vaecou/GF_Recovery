package cmd

import (
	"context"
	"time"

	"GF_Recovery/internal/consts"
	"GF_Recovery/internal/controller/admin"
	"GF_Recovery/internal/middleware"
	"GF_Recovery/utility/simple"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
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

			// 注册全局中间件
			s.BindMiddleware("/*", []ghttp.HandlerFunc{
				middleware.CORS,                 // 跨域中间件，自动处理跨域问题
				ghttp.MiddlewareHandlerResponse, // HTTP响应预处理，在业务处理完成后，对响应结果进行格式化和错误过滤，将处理后的数据发送给请求方
			}...)

			s.Group("/admin", func(group *ghttp.RouterGroup) {
				// 注册分组中间件
				middleware.AdiminUserToken(ctx, group) // Token以及登录中间件
				group.Map(g.Map{
					// 增加管理员
					"POST:/user": admin.AdminUser.AddAdminUser,
					// 修改管理员
					"PUT:/user": admin.AdminUser.UpdateAdminUser,
					// 查询管理员列表
					"GET:/user": admin.AdminUser.GetAdminUser,
					// 查询管理员列表
					"GET:/user/list": admin.AdminUser.GetAdminUserList,
				})
				group.Map(g.Map{
					// 增加设置
					"POST:/about": admin.Setting.AddOrUpdateSetting,
					// 修改设置
					"PUT:/about": admin.Setting.AddOrUpdateSetting,
					// 查询设置
					"GET:/about": admin.Setting.GetSetting,
				})
				group.Map(g.Map{
					// 增加问题
					"POST:/question": admin.Question.AddQuestion,
					// 修改问题
					"PUT:/question": admin.Question.UpdateQuestion,
					// 删除问题
					"DELETE:/question": admin.Question.DeleteQuestion,
					// 查询问题
					"GET:/question": admin.Question.GetQuestion,
					// 查询问题列表
					"GET:/question/list": admin.Question.GetQuestionList,
				})
			})

			serverWg.Add(1)
			// 信号监听
			signalListen(ctx, signalHandlerForOverall)
			go func() {
				<-serverCloseSignal
				g.Log().Debug(ctx, "Http已成功关闭")
				serverWg.Done()
			}()

			s.Run()

			return
		},
	}
)
