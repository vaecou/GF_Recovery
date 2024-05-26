package middleware

import (
	"GF_Recovery/internal/consts"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/model/do"
	"GF_Recovery/internal/model/entity"
	"GF_Recovery/utility/validate"
	"context"
	"time"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
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

func MiniUserToken(ctx context.Context, group *ghttp.RouterGroup) {
	// 启动gtoken
	gtoken := &gtoken.GfToken{
		LoginPath:        "POST:/login",
		LoginBeforeFunc:  miniUserLoginBefore,
		LoginAfterFunc:   miniUserLoginAfter,
		LogoutPath:       "DELETE:/login",
		AuthAfterFunc:    miniUserAuthAfter,
		AuthExcludePaths: g.SliceStr{"/mini/login"},
		CacheMode:        2,
		CacheKey:         "Mini:User:Login:User_ID_",
		Timeout:          31536000000,
		EncryptKey:       gconv.UnsafeStrToBytes("75315978932147857291236985632147"),
	}
	err := gtoken.Middleware(ctx, group)
	if err != nil {
		panic(err)
	}
}

func miniUserLoginBefore(r *ghttp.Request) (id string, data interface{}) {
	ctx := r.GetCtx()
	code := r.GetForm("code").String()
	if code == "" {
		r.Response.WriteJson(DefaultHandlerResponse{
			Code:    gcode.CodeValidationFailed.Code(),
			Message: "Code不能为空",
		})
		return
	}

	appid := g.Cfg().MustGet(ctx, "miniprogram.appid").String()
	secret := g.Cfg().MustGet(ctx, "miniprogram.secret").String()
	addrs := g.Cfg().MustGet(ctx, "redis.default.address").String()
	MiniProgramApp, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:     appid,  // 小程序appid
		Secret:    secret, // 小程序app secret
		HttpDebug: true,
		// 可选，不传默认走程序内存
		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{addrs},
			Password: "",
			DB:       0,
		}),
	})
	if err != nil {
		panic(err)
	}

	res, err := MiniProgramApp.Auth.Session(ctx, code)
	if err != nil {
		panic(err)
	}

	var mb *entity.ReAdminUser
	if err := dao.ReUser.Ctx(ctx).Where("unionid", res.UnionID).Scan(&mb); err != nil {
		r.Response.WriteJson(DefaultHandlerResponse{
			Code:    gcode.CodeInternalError.Code(),
			Message: consts.ErrorORM,
		})
		return
	}
	var result int64
	// 不存在unionid
	if mb == nil {
		user := &do.ReUser{
			Openid:  res.OpenID,
			Unionid: res.UnionID,
			Status:  0,
		}
		res, err := dao.ReUser.Ctx(ctx).Insert(user)
		if err != nil {
			panic(err)
		}
		result, err = res.LastInsertId()
		if err != nil {
			panic(err)
		}
	} else {
		result = gconv.Int64(mb.Id)
	}

	id = gconv.String(result)
	data = g.Map{
		"openid":  res.OpenID,
		"unionid": res.UnionID,
	}
	return
}

func miniUserLoginAfter(r *ghttp.Request, respData gtoken.Resp) {
	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)

	if respData.Success() {
		code = gcode.CodeOK
		res = g.Map{
			"token": respData.GetString("token"),
		}
	}

	r.Response.WriteJson(DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}

// 自定义验证返回的格式
func miniUserAuthAfter(r *ghttp.Request, respData gtoken.Resp) {
	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)

	if respData.Success() {
		r.SetCtxVar("user_id", respData.Get("userKey"))
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
