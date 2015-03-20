package main

import (
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/yuhaya/CloudServer/controller"
	"github.com/yuhaya/CloudServer/models"
	"github.com/yuhaya/CloudServer/tool"
	"log"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
	CreatTables()
}

func NotFound(w http.ResponseWriter, r *http.Request) {

}

func main() {
	m := pat.New()
	m.Get("/", http.HandlerFunc(Root))
	//数据查询示例
	m.Get("/sql", http.HandlerFunc(controller.SqlDemo))

	err := http.ListenAndServe(":9003", m)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func CreatTables() {
	db_gorm := tool.GetGormDB()
	db_gorm.CreateTable(&models.Admins{})
}
