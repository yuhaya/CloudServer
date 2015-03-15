package tool

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/yuhaya/CloudServer/conf"
)

var db_connect *sql.DB

func init() {
	db, err := gorm.Open("mysql", conf.GlobalConfig["mysql"])
	if err == nil {
		db_connect = db.DB()
		fmt.Print("Database Connect Success !")
	} else {
		fmt.Print("Database Connect Faile !")
		panic(err)
	}
}

func GetDbConnect() *sql.DB {
	return db_connect
}
