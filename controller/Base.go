package controller

import (
	"net/http"
)

type ControllerInterface interface {
	Before(ResponseWriter http.ResponseWriter, Request *http.Request, ActionName string, ControllerName string)
	Pre()
	End()
}

type Base struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	ActionName     string
	ControllerName string
}

func (this *Base) Before(ResponseWriter http.ResponseWriter, Request *http.Request, ActionName string, ControllerName string) {
	this.ResponseWriter = ResponseWriter
	this.Request = Request
	this.ActionName = ActionName
	this.ControllerName = ControllerName
}
func (this *Base) Pre() {

}
func (this *Base) End() {

}

func SqlDemo(w http.ResponseWriter, r *http.Request) {

	//	data := models.Query("SELECT * FROM ittr_cards")
	//	fmt.Fprintf(w, "\n%#v\n", data)
}
