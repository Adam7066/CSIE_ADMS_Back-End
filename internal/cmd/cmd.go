package cmd

import (
	"CSIE_ADMS_Back-End/internal/controller/admsInfo"
	"CSIE_ADMS_Back-End/internal/controller/auth"
	"CSIE_ADMS_Back-End/internal/controller/hello"
	"CSIE_ADMS_Back-End/internal/controller/rank"
	"CSIE_ADMS_Back-End/internal/controller/user"
	"CSIE_ADMS_Back-End/utility"
	"context"
	"database/sql"
	"fmt"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			gfToken := &gtoken.GfToken{
				CacheMode:       CfgGet(ctx, "gToken.CacheMode").Int8(),
				MultiLogin:      CfgGet(ctx, "gToken.MultiLogin").Bool(),
				Timeout:         CfgGet(ctx, "gToken.Timeout").Int(),
				EncryptKey:      CfgGet(ctx, "gToken.EncryptKey").Bytes(),
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
					auth.NewV1(),
					admsInfo.NewV1(),
					rank.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)

func CfgGet(ctx context.Context, name string) *gvar.Var {
	gVa, _ := g.Config().Get(ctx, name)
	return gVa
}

func MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"localhost"}
	if !r.Response.CORSAllowedOrigin(corsOptions) {
		r.Response.WriteStatusExit(http.StatusForbidden)
	}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

func InitDB() {
	// remove old db
	err := os.Remove("./CSIE_ADMS_DB.sqlite3")
	utility.IfErrExit(err)

	// init db
	db, err := sql.Open("sqlite3", "./CSIE_ADMS_DB.sqlite3")
	utility.IfErrExit(err)
	defer db.Close()

	// init tables
	sqlFilename := []string{
		"users.sql", "students.sql", "adms.sql", "rank.sql",
	}
	for _, filename := range sqlFilename {
		sqlFile, err := os.ReadFile("./manifest/sql/" + filename)
		utility.IfErrExit(err)
		_, err = db.Exec(string(sqlFile))
		utility.IfErrExit(err)
	}
	fmt.Println("init db & tables success")

	// insert default data
	initTablesData()
	fmt.Println("insert default data success")
}
