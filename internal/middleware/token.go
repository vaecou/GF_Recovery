package middleware

import (
	"GF_Recovery/internal/consts"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/model/entity"
	"GF_Recovery/utility/validate"
	"context"
	"fmt"
	"time"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type DefaultHandlerResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

func AdiminUserToken(ctx context.Context, group *ghttp.RouterGroup) {
	// 启动gtoken
	gtoken := &gtoken.GfToken{
		LoginPath:        "POST:/login",
		LoginBeforeFunc:  adminUserLoginBefore,
		LoginAfterFunc:   adminUserLoginAfter,
		LogoutPath:       "DELETE:/login",
		AuthAfterFunc:    adminUserAuthAfter,
		AuthExcludePaths: g.SliceStr{"/admin/login"},
		CacheMode:        2,
		CacheKey:         "Admin:User:Login:User_ID_",
		Timeout:          gconv.Int(time.Hour.Milliseconds()),
		EncryptKey:       gconv.UnsafeStrToBytes("75315978932147896541236985632147"),
	}
	err := gtoken.Middleware(ctx, group)
	if err != nil {
		panic(err)
	}
}

// 自定义登录前校验
func adminUserLoginBefore(r *ghttp.Request) (id string, data interface{}) {
	ctx := r.GetCtx()
	account := r.GetForm("account").String()
	password := r.GetForm("password").String()

	if account == "" {
		r.Response.WriteJson(DefaultHandlerResponse{
			Code:    gcode.CodeValidationFailed.Code(),
			Message: "请输入账号",
		})
		return
	}

	if password == "" {
		r.Response.WriteJson(DefaultHandlerResponse{
			Code:    gcode.CodeValidationFailed.Code(),
			Message: "密码不能为空",
		})

		return
	}

	if len(password) < 6 || len(password) > 18 {
		r.Response.WriteJson(DefaultHandlerResponse{
			Code:    gcode.CodeValidationFailed.Code(),
			Message: "密码长度在6-18位之间",
		})

		return
	}

	var mb *entity.ReAdminUser

	// 判断Account是否存在
	if err := dao.ReAdminUser.Ctx(ctx).Where("account", account).Scan(&mb); err != nil {
		r.Response.WriteJson(DefaultHandlerResponse{
			Code:    gcode.CodeInternalError.Code(),
			Message: consts.ErrorORM,
		})
		return
	}

	if mb == nil {
		r.Response.WriteJson(DefaultHandlerResponse{
			Code:    gcode.CodeInternalError.Code(),
			Message: "账号不存在",
		})
		return
	}

	// 验证密码
	if err := validate.CheckPassword(password, mb.Salt, mb.Password); err != nil {
		r.Response.WriteJson(DefaultHandlerResponse{
			Code:    gcode.CodeInternalError.Code(),
			Message: err.Error(),
		})

		return
	}

	// 记录登录信息
	id = gconv.String(mb.Id)
	data = g.Map{
		"name":       mb.Name,
		"account":    mb.Account,
		"created_at": mb.CreatedAt,
	}

	// 将name存储至ctx中
	r.SetCtxVar("name", mb.Name)
	// 将account存储至ctx中
	r.SetCtxVar("account", mb.Account)

	return
}

// 自定义登录后的记录
func adminUserLoginAfter(r *ghttp.Request, respData gtoken.Resp) {
	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)

	if respData.Success() {
		code = gcode.CodeOK
		fmt.Println("respData", respData)
		res = g.Map{
			"name":    r.GetCtxVar("name"),
			"account": r.GetCtxVar("account"),
			"token":   respData.GetString("token"),
		}
	}

	r.Response.WriteJson(DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}

// 自定义验证返回的格式
func adminUserAuthAfter(r *ghttp.Request, respData gtoken.Resp) {
	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)

	if respData.Success() {
		r.Middleware.Next()
		return
	} else {
		code = gcode.CodeNotAuthorized
		msg = "Token验证失败"
	}

	r.Response.WriteJson(DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}
