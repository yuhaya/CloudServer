package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yuhaya/CloudServer/tool"
)

func Query(sqlstring string, vals ...interface{}) []map[string]string {
	db := tool.GetDbConnect()

	var rows *sql.Rows
	var err error

	if len(vals) > 0 {
		rows, err = db.Query(sqlstring, vals)
	} else {
		rows, err = db.Query(sqlstring)
	}

	if err != nil {
		fmt.Print(err.Error())
	}

	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	result := make([]map[string]string, 0, 10000)
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
		result = append(result, record)
	}
	fmt.Printf("\n%#v\n", result)
	return result
}
