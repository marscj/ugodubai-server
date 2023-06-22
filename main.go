package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "ugodubai-server/internal/app/boot"

	"github.com/gogf/gf/v2/os/gctx"
	_ "ugodubai-server/internal/app/system/packed"
	"ugodubai-server/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
