package main

import (
	_ "CSIE_ADMS_Back-End/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	"database/sql"
	"flag"
	"fmt"
	"os"

	"CSIE_ADMS_Back-End/internal/cmd"
	"CSIE_ADMS_Back-End/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	initFlag := flag.Bool("init", false, "init")
	flag.Parse()
	if *initFlag {
		initDB()
	} else {
		cmd.Main.Run(gctx.GetInitCtx())
	}
}

func initDB() {
	// init db
	db, err := sql.Open("sqlite3", "./CSIE_ADMS_DB.sqlite3")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()
	// init tables
	sqlFilename := []string{
		"users.sql", "students.sql", "adms.sql",
	}
	for _, filename := range sqlFilename {
		sqlFile, err := os.ReadFile("./manifest/sql/" + filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = db.Exec(string(sqlFile))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("init db success")

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
	fmt.Println("insert default admin success")
}
