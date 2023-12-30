package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.Create("version", "description")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	table.AddRow([]string{"gotable 5", "Safe: New table type to enhance concurrency security"})
	table.AddRow([]string{"gotab我是中国&a2语言测试le 4？", "Colored: Print colored column"})
	table.AddRow([]string{"gotabl牛逼？e 3", "Storage: Store the table data as a file"})
	table.AddRow([]string{"gotab我1是中le 2", "Simple: Use simpler APIs to control table"})
	table.AddRow([]string{"gotable 1", "Gotable: Print a beautiful ASCII table"})

	fmt.Println(table)
}
