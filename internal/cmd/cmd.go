package cmd

import (
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
	"net/http"
	"os"

	"CSIE_ADMS_Back-End/internal/controller/hello"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			initDB(ctx)

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

func initDB(ctx context.Context) {
	// check db exist
	dbName := CfgGet(ctx, "database.default.filename").String()
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// check table exist
	if exist, err := isTableExist(db, "users"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else if !exist {
		// create users table
		sqlFile, err := os.ReadFile("./manifest/sql/users.sql")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = db.Exec(string(sqlFile))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// insert default admin
		m := g.DB().Model("users")
		encodedHash, _ := utility.HashPWD("1234567890")
		_, err = m.Insert(g.Map{
			"email":    "admin@csie.adms",
			"password": encodedHash,
			"username": "admin",
			"role":     "admin",
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func isTableExist(db *sql.DB, tableName string) (bool, error) {
	query := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s';", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return rows.Next(), nil
}
