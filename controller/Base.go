// @APIVersion 1.0.0
// @Title demo API
// @Description ${PROJECT_NAME} | this is demo description
// @Contact mao | 3wmaocomputer@gmail.com
// @Date 3/12/15

package controller

import (
	"net/http"
)

type Base struct {
	ResponseWriter *http.ResponseWriter
	Request        *http.Request
}
