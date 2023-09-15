package cmd

import (
	"CSIE_ADMS_Back-End/internal/controller/user"
	"context"
	"github.com/goflyfox/gtoken/gtoken"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"CSIE_ADMS_Back-End/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			gfToken := &gtoken.GfToken{
				LoginPath:       "/login",
				LoginBeforeFunc: loginFunc,
				LoginAfterFunc:  loginAfterFunc,
				LogoutPath:      "/logout",
			}

			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Bind(
					hello.NewV1(),
					user.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
