// @APIVersion 1.0.0
// @Title demo API
// @Description ${PROJECT_NAME} | this is demo description
// @Contact mao | 3wmaocomputer@gmail.com
// @Date 3/12/15

package controller

import (
	"fmt"
	"github.com/yuhaya/CloudServer/model"
	"net/http"
)

func SqlDemo(w http.ResponseWriter, r *http.Request) {

	data := model.Query("SELECT * FROM ittr_cards")
	fmt.Fprintf(w, "\n%#v\n", data)
}
