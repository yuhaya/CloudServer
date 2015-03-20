package tool

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/yuhaya/CloudServer/conf"
)

var db_connect *sql.DB
var db_gorm *gorm.DB

func init() {
	db, err := gorm.Open("mysql", conf.GlobalConfig["mysql"])
	if err == nil {
		db_gorm = &db
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

func GetGormDB() *gorm.DB {
	return db_gorm
}
