package main

import (
	"fmt"
	"github.com/foolin/sumdiff/internal/app"
	"github.com/foolin/sumdiff/internal/plog"
)

func main() {
	ok, err := app.DiffDir("../../test_data/data4", "../../test_data/data3")
	if err != nil {
		plog.PrintProgress(err.Error())
	}
	plog.Print(fmt.Sprintf("%v", ok))
}
