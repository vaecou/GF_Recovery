package mini

import (
	"GF_Recovery/api/mini"
	"GF_Recovery/internal/dao"
	"GF_Recovery/internal/model/entity"
	"GF_Recovery/internal/service"
	"GF_Recovery/utility/generate"
	"context"
	"fmt"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/transfer/request"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sMiniBalance struct {
}

func init() {
	service.RegisterMiniBalance(NewBalance())
}

func NewBalance() *sMiniBalance {
	return &sMiniBalance{}
}

func (c *sMiniBalance) GetUserBalance(ctx context.Context, req *mini.GetBalanceInfoReq) (res []*mini.BalanceListRes, err error) {
	id := gconv.Int(ctx.Value("user_id"))
	if req.Type == 0 {
		err = dao.ReTakeList.Ctx(ctx).Where("user_id", id).Page(req.Page, req.Limit).OrderDesc("created_at").Scan(&res)
		if err != nil {
			return
		}
	} else {
		err = dao.ReTakeList.Ctx(ctx).Where("user_id", id).Where("type", req.Type).Page(req.Page, req.Limit).OrderDesc("created_at").Scan(&res)
		if err != nil {
			return
		}
	}
	return
}

// 提现
func (c *sMiniBalance) Withdraw(ctx context.Context, req *mini.WithdrawReq) (err error) {
	// 开启事务
	dao.ReTakeList.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		// 插入一个提现记录
		id := gconv.Int(ctx.Value("user_id"))
		order_id, err := generate.GenerateOrderNo("0")
		if err != nil {
			return err
		}
		_, err = dao.ReTakeList.Ctx(ctx).TX(tx).Insert(g.Map{
			"id":      order_id,
			"user_id": id,
			"balance": req.Money,
			"type":    2,
		})
		if err != nil {
			return err
		}

		// 减少用户的余额
		_, err = dao.ReUser.Ctx(ctx).TX(tx).Where("id", id).Decrement("balance", req.Money)
		if err != nil {
			return err
		}

		appid := g.Cfg().MustGet(ctx, "payment.appid").String()
		mchid := g.Cfg().MustGet(ctx, "payment.mchid").String()
		mchapiv3key := g.Cfg().MustGet(ctx, "payment.mchapiv3key").String()
		certpath := g.Cfg().MustGet(ctx, "payment.certpath").String()
		keypath := g.Cfg().MustGet(ctx, "payment.keypath").String()
		serialno := g.Cfg().MustGet(ctx, "payment.serialno").String()
		addrs := g.Cfg().MustGet(ctx, "redis.default.address").String()

		fmt.Println("appid", appid)
		fmt.Println("mchid", mchid)
		fmt.Println("mchapiv3key", mchapiv3key)
		fmt.Println("certpath", certpath)
		fmt.Println("keypath", keypath)
		fmt.Println("serialno", serialno)
		fmt.Println("addrs", addrs)

		PaymentService, err := payment.NewPayment(&payment.UserConfig{
			AppID:       appid,       // 小程序、公众号或者企业微信的appid
			MchID:       mchid,       // 商户号 appID
			MchApiV3Key: mchapiv3key, // 微信V3接口调用必填
			CertPath:    certpath,    // 商户后台支付的Cert证书路径
			KeyPath:     keypath,     // 商户后台支付的Key证书路径
			SerialNo:    serialno,    // 商户支付证书序列号
			// NotifyURL:   "[notify_url]",
			HttpDebug: true,
			// 可选，不传默认走程序内存
			Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
				Addrs:    []string{addrs},
				Password: "",
				DB:       0,
			}),
		})
		if err != nil {
			return err
		}

		// 获取openid
		user := entity.ReUser{}
		dao.ReUser.Ctx(ctx).TX(tx).Where("id", id).Scan(&user)

		fmt.Println("PartnerTradeNo", gconv.String(order_id))
		fmt.Println("OpenID", user.Openid)
		fmt.Println("Amount", req.Money*100)

		transfer := &request.RequestTransferBatch{
			AppID:       appid,
			OutBatchNO:  gconv.String(order_id),
			BatchName:   "余额提现",
			BatchRemark: "余额提现",
			TotalAmount: req.Money * 100,
			TotalNum:    1,
			TransferDetailList: []*request.TransferDetail{
				{
					OutDetailNO:    gconv.String(order_id),
					TransferAmount: req.Money * 100,
					TransferRemark: "旧衣回收提现",
					OpenID:         user.Openid,
				},
			},
		}

		payResult, err := PaymentService.TransferBatch.Batch(ctx, transfer)
		if err != nil {
			panic(err)
		}

		fmt.Println("payResult", payResult)

		return nil
	})
	return
}
