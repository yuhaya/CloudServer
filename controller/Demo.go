// @APIVersion 1.0.0
// @Title demo API
// @Description ${PROJECT_NAME} | this is demo description
// @Contact mao | 3wmaocomputer@gmail.com
// @Date 3/12/15

package controller

import (
	"fmt"
	"github.com/yuhaya/CloudServer/model"
	//	"github.com/yuhaya/CloudServer/tool"
	"net/http"
)

func SqlDemo(w http.ResponseWriter, r *http.Request) {
	//	db := tool.GetDbConnect()
	//	rows, _ := db.Query("SELECT * FROM demo")

	//	for rows.Next() {
	//		var userId int
	//		var userName string
	//		rows.Columns()
	//		rows.Scan(&userId, &userName)

	//		fmt.Fprint(w, userId)
	//		fmt.Fprint(w, userName)
	//	}
	//	fmt.Fprintf(w, "<h1>%s</h1>", "DemoSql2")

	data := model.Query("SELECT * FROM demo")
	fmt.Fprintf(w, "\n%#v\n", data)
}
