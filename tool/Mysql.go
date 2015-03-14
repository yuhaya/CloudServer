package tool

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yuhaya/CloudServer/conf"
)

var db_connect *sql.DB

func init() {
	db, err := sql.Open("mysql", conf.GlobalConfig["mysql"])
	if err == nil {
		db_connect = db
		fmt.Print("Database Connect Success !")
	} else {
		fmt.Print("Database Connect Faile !")
		panic(err)
	}
}

func GetDbConnect() *sql.DB {
	return db_connect
}
