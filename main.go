package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"m2f/root"
	files "m2f/utils"
)

func main() {
	instance := root.InstanceDB()
	tables := instance.GetTables()
	fmt.Printf("扫描到表 %d 张\n", len(tables))
	for i := 0; i < len(tables); i++ {
		fmt.Printf("%d-%s\n", i+1, tables[i])
	}
	fmt.Println("开始处理表结构")
	maps := root.GetMaps()
	for _, table := range tables {
		m2ms := instance.GetTableSchema(table)
		var str string
		var firstRow = true
		for _, m2m := range m2ms {
			columns := files.StrToUpper(m2m.ColumnName)
			if firstRow {
				str += columns + " " + maps[m2m.DataType] + "\n"
				firstRow = false
			} else {
				str += "	" + columns + " " + maps[m2m.DataType] + "\n"
			}
		}
		table := files.StrToUpper(table)
		ok, err := files.GenerateStructs("demo/"+table+".go", table, str)
		if ok {
			fmt.Printf("%s表创建成功\n", table)
		} else {
			fmt.Printf("%s表创建失败\n", err.Error())
		}
	}
}
