package controller

import (
	//	"fmt"
	"github.com/bmizerany/pat"
	"net/http"
	"strconv"
)

type NextPreer interface {
	NextPre()
}

type ControllerInterface interface {
	Init(ResponseWriter http.ResponseWriter, Request *http.Request, ActionName string, ControllerName string, AppController interface{}, Router *pat.PatternServeMux)
	Pre()
	End()
}

type Base struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	ActionName     string
	ControllerName string
	AppController  interface{}
	Router         *pat.PatternServeMux
}

func (this *Base) Init(ResponseWriter http.ResponseWriter, Request *http.Request, ActionName string, ControllerName string, AppController interface{}, Router *pat.PatternServeMux) {
	this.ResponseWriter = ResponseWriter
	this.Request = Request
	this.ActionName = ActionName
	this.ControllerName = ControllerName
	this.AppController = AppController
	this.Router = Router
	this.Pre()
}
func (this *Base) Pre() {
	if app, ok := this.AppController.(NextPreer); ok {
		app.NextPre()
	}
}
func (this *Base) End() {
	//	fmt.Fprint(this.ResponseWriter, "End")
}

func (this *Base) GetIntUrlParam(v ...interface{}) (int, bool) {
	par_num := len(v)
	if par_num < 1 {

		return 0, false

	} else if par_num == 1 {

		if val_key, ok := v[0].(string); ok {
			val := this.Request.URL.Query().Get(val_key)
			if val_end, err := strconv.Atoi(val); err == nil {
				return val_end, true
			}
		}
		return 0, false

	} else {

		var par_default int

		if val_default, ok := v[1].(int); ok {
			par_default = val_default
		} else {
			par_default = 0
		}

		if val_key, ok := v[0].(string); ok {
			val := this.Request.URL.Query().Get(val_key)
			if val_end, err := strconv.Atoi(val); err == nil {
				return val_end, true
			}
		}
		return par_default, true
	}
}

func (this *Base) GetStringUrlParam(v ...interface{}) string {

	par_num := len(v)
	if par_num < 1 {
		return ""

	} else if par_num == 1 {

		if val_key, ok := v[0].(string); ok {

			return this.Request.URL.Query().Get(val_key)

		} else {
			//参数不存在
			return ""
		}

	} else {

		var par_default string

		if val_default, ok := v[1].(string); ok {
			par_default = val_default
		} else {
			par_default = ""
		}

		if val_key, ok := v[0].(string); ok {

			return this.Request.URL.Query().Get(val_key)

		} else {
			//参数不存在
			return par_default
		}
	}
}

func (this *Base) GetIntPostParam(v ...interface{}) (int, bool) {

	par_num := len(v)
	if par_num < 1 {

		return 0, false

	} else if par_num == 1 {

		if val_key, ok := v[0].(string); ok {

			this.Request.ParseForm()
			if len(this.Request.PostForm[val_key]) > 0 {
				if val_end, err := strconv.Atoi(this.Request.PostForm[val_key][0]); err == nil {
					return val_end, true
				}
			}
		}
		return 0, false

	} else {

		var par_default int

		if val_default, ok := v[1].(int); ok {
			par_default = val_default
		} else {
			par_default = 0
		}

		if val_key, ok := v[0].(string); ok {

			this.Request.ParseForm()
			if len(this.Request.PostForm[val_key]) > 0 {
				if val_end, err := strconv.Atoi(this.Request.PostForm[val_key][0]); err == nil {
					return val_end, true
				}
			}

		}
		return par_default, true
	}
}

func (this *Base) GetStringPostParam(v ...interface{}) string {

	par_num := len(v)
	if par_num < 1 {

		return ""

	} else if par_num == 1 {

		if val_key, ok := v[0].(string); ok {

			this.Request.ParseForm()
			if len(this.Request.PostForm[val_key]) > 0 {
				return this.Request.PostForm[val_key][0]
			}
		}
		return ""

	} else {

		var par_default string

		if val_default, ok := v[1].(string); ok {
			par_default = val_default
		} else {
			par_default = ""
		}

		if val_key, ok := v[0].(string); ok {

			this.Request.ParseForm()
			if len(this.Request.PostForm[val_key]) > 0 {
				return this.Request.PostForm[val_key][0]
			}

		}
		return par_default
	}
}

func SqlDemo(w http.ResponseWriter, r *http.Request) {

	//	data := models.Query("SELECT * FROM ittr_cards")
	//	fmt.Fprintf(w, "\n%#v\n", data)
}
