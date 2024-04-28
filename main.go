package main

import (
	"GF_Recovery/internal/boot"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"GF_Recovery/internal/cmd"
)

func main() {
	ctx := gctx.GetInitCtx()
	boot.Init(ctx)
	cmd.Main.Run(ctx)
}
