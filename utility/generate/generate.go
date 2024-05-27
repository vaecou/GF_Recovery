package generate

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

// 生成订单编号函数
func GenerateOrderNo(sourceType string) (int64, error) {
	// 获取当前时间
	currentTime := time.Now()

	// 格式化当前时间为【年的后2位+月+日】
	originDateStr := currentTime.Format("060102")

	// 计算当前时间走过的秒
	startOfDay := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	differSecond := int(currentTime.Sub(startOfDay).Seconds())

	// 获取【年的后2位+月+日+秒】，秒的长度不足补充0
	yyMMddSecond := originDateStr + fmt.Sprintf("%05d", differSecond)

	// 获取【业务编码】 + 【年的后2位+月+日+秒】，作为自增key；
	prefixOrder := sourceType + yyMMddSecond

	// 通过key，采用redis自增函数，实现单秒自增；不同的key，从0开始自增，同时设置60秒过期
	ctx := gctx.New()
	incrId, err := g.Redis().Incr(ctx, prefixOrder)
	if err != nil {
		return 0, err
	}

	// 设置key过期时间为60秒
	_, err = g.Redis().Expire(ctx, prefixOrder, 60)
	if err != nil {
		return 0, err
	}

	// 生成订单编号
	orderNo := gconv.Int64(prefixOrder + fmt.Sprintf("%02d", incrId))

	return orderNo, nil
}
