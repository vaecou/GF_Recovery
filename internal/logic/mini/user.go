package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/model/do"
	"GF_Recovery/internal/model/entity"
	"GF_Recovery/internal/service"
	"context"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sMiniUsers struct {
}

func init() {
	service.RegisterMiniUsers(NewMiniUsers())
}

func NewMiniUsers() *sMiniUsers {
	return &sMiniUsers{}
}

// 检查用户是否存在phone
func (s *sMiniUsers) CheckUserPhone(ctx context.Context) (res bool, err error) {
	id := gconv.Int(ctx.Value("user_id"))
	user := do.ReUser{}
	err = dao.ReUser.Ctx(ctx).Where("id", id).Scan(&user)
	if err != nil {
		return
	}
	if user.Phone != nil {
		res = true
	}
	return
}

func (s *sMiniUsers) SaveUserPhone(ctx context.Context, req *mini.SaveUserPhoneReq) (res *mini.UserPhoneRes, err error) {
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

	// {"errcode":0,"errmsg":"ok","phone_info":{"phoneNumber":"19916545450","purePhoneNumber":"19916545450","countryCode":"86","watermark":{"timestamp":1716818136,"appid":"wx3f7c1634de736827"}}}
	result, err := MiniProgramApp.PhoneNumber.GetUserPhoneNumber(ctx, req.Code)
	if err != nil {
		return
	}

	id := gconv.Int(ctx.Value("user_id"))

	var mb *entity.ReUser
	err = dao.ReUser.Ctx(ctx).Where("id", id).Scan(&mb)
	if err != nil {
		return
	}

	// 添加用户的手机号
	_, err = dao.ReUser.Ctx(ctx).OmitEmpty().Where("id", id).Update(g.Map{
		"phone": result.PhoneInfo.PhoneNumber,
	})
	if err != nil {
		return
	}
	return
}
