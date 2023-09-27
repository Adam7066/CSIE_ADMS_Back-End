package cmd

import (
	"CSIE_ADMS_Back-End/internal/controller/user"
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"net/http"

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
				AuthAfterFunc:   authAfterFunc,
			}

			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(MiddlewareCORS)
				group.Bind(
					hello.NewV1(),
				)

				err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Bind(
					user.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)

func MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"localhost"}
	if !r.Response.CORSAllowedOrigin(corsOptions) {
		r.Response.WriteStatusExit(http.StatusForbidden)
	}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
