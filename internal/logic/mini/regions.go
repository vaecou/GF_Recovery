package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sMiniRegions struct {
}

func init() {
	service.RegisterMiniRegions(NewRegions())
}

func NewRegions() *sMiniRegions {
	return &sMiniRegions{}
}

// 获取地址
func (s *sMiniRegions) GetAddress(ctx context.Context, in *mini.GetAddressReq) (res *mini.AddressRes, err error) {
	err = dao.ReAddressList.Ctx(ctx).Where("id", in.ID).Scan(&res)
	if err != nil {
		return
	}
	return
}

func (s *sMiniRegions) DeleteAddress(ctx context.Context, in *mini.DeleteAddressReq) (err error) {
	_, err = dao.ReAddressList.Ctx(ctx).Where("id", in.ID).Delete()
	if err != nil {
		return
	}
	return
}

func (s *sMiniRegions) UpdateAddress(ctx context.Context, in *mini.UpdateAddressReq) (err error) {
	_, err = dao.ReAddressList.Ctx(ctx).Where("id", in.ID).Update(in)
	if err != nil {
		return
	}
	return
}

func (s *sMiniRegions) UpdateRadio(ctx context.Context, in *mini.UpdateRadioReq) (err error) {
	id := ctx.Value("user_id")
	// 将user_id下的default设置为false
	_, err = dao.ReAddressList.Ctx(ctx).Where("user_id", id).Update(g.Map{
		"default": false,
	})
	if err != nil {
		return
	}
	_, err = dao.ReAddressList.Ctx(ctx).Where("id", in.ID).Update(g.Map{
		"default": true,
	})
	return
}

func (s *sMiniRegions) AddAddress(ctx context.Context, in *mini.AddRegionsReq) (err error) {
	id := ctx.Value("user_id")
	in.UserID = gconv.Int(id)
	if in.Default {
		// 将其它的default设置为false
		_, err = dao.ReAddressList.Ctx(ctx).Where("user_id", id).Update(g.Map{
			"default": false,
		})
		if err != nil {
			return
		}
	}
	_, err = dao.ReAddressList.Ctx(ctx).Insert(in)
	if err != nil {
		return
	}
	return
}

func (c *sMiniRegions) GetAddressRadio(ctx context.Context, id int) (res int, err error) {
	address := &mini.AddressRes{}
	err = dao.ReAddressList.Ctx(ctx).Where("user_id = ? AND `default` = ?", id, true).Scan(&address)
	if err != nil {
		return
	}
	res = address.ID
	return
}

func (c *sMiniRegions) GetAddressList(ctx context.Context, id int) (res []*mini.AddressRes, err error) {
	err = dao.ReAddressList.Ctx(ctx).Where("user_id", id).OrderDesc("default").Scan(&res)
	if err != nil {
		return
	}
	return
}
