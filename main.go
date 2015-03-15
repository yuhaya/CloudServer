package main

import (
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/yuhaya/CloudServer/controller"
	"log"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func NotFound(w http.ResponseWriter, r *http.Request) {

}

func main() {
	m := pat.New()
	m.Get("/", http.HandlerFunc(Root))
	//数据查询示例
	m.Get("/sql", http.HandlerFunc(controller.SqlDemo))

	err := http.ListenAndServe(":9000", m)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
