package main

import (
	"CSIE_ADMS_Back-End/internal/cmd"
	_ "CSIE_ADMS_Back-End/internal/packed"
	"flag"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	initFlag := flag.Bool("init", false, "init")
	flag.Parse()
	if *initFlag {
		cmd.InitDB()
	} else {
		cmd.Main.Run(gctx.GetInitCtx())
	}
}
