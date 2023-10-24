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
	utility.IfErrExit(err)
	defer db.Close()

	// init tables
	sqlFilename := []string{
		"users.sql", "students.sql", "adms.sql",
	}
	for _, filename := range sqlFilename {
		sqlFile, err := os.ReadFile("./manifest/sql/" + filename)
		utility.IfErrExit(err)
		_, err = db.Exec(string(sqlFile))
		utility.IfErrExit(err)
	}
	fmt.Println("init db & tables success")

	// insert default data
	cmd.InitTablesData()

	fmt.Println("insert default data success")
}
