package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/util/gconv"
)

var Personal = cPersonal{}

type cPersonal struct {
}

func (c *cPersonal) GetInfo(ctx context.Context, req *mini.GetInfoReq) (res *mini.InfoRes, err error) {
	res = &mini.InfoRes{}
	kg, err := service.MiniOrder().GetUserKG(ctx)
	if err != nil {
		return
	}
	balance, err := service.MiniUsers().GetUserBalance(ctx)
	if err != nil {
		return
	}
	res.KG = gconv.Float64(fmt.Sprintf("%.3f", kg))
	res.Balance = balance
	res.One = gconv.Float64(fmt.Sprintf("%.3f", res.KG*0.19))
	res.Two = gconv.Float64(fmt.Sprintf("%.3f", res.KG*3.6))
	res.Three = gconv.Float64(fmt.Sprintf("%.3f", res.KG*0.92))
	return
}
