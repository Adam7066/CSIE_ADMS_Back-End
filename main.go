package main

import (
	_ "CSIE_ADMS_Back-End/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"CSIE_ADMS_Back-End/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
