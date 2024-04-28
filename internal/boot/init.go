package boot

import (
	"GF_Recovery/internal/boot/mode"
	"GF_Recovery/internal/boot/time"
	"context"
	"fmt"
	"runtime"

	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/os/gfile"
)

func Init(ctx context.Context) {
	// 设置运行模式
	mode.SetGFMode(ctx)

	// 设置时区
	time.SetTimeZone(ctx)

	fmt.Printf("\r\nGoFrame版本：%v\r\n当前运行环境：%v\r\n运行根路径为：%v\r\n\r\n", gf.VERSION, runtime.GOOS, gfile.Pwd())
}
