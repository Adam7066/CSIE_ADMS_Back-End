package utility

import (
	"fmt"
	"os"
)

func IfErrExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
