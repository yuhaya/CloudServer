package controller

import (
	"fmt"
)

type User struct {
	Base
}

func (this User) NextPre() {
	//	fmt.Fprint(this.ResponseWriter, "NextPre")
}

func (this User) ActionHello() {
	fmt.Fprint(this.ResponseWriter, "Hello Go !")
}

//用户创建
func (this User) ActionCreate() {
	this.Request.ParseForm()
	fmt.Fprintf(this.ResponseWriter, "\nxx%#vxx\n", this.Request.Form)

	//	card := this.Request.URL.Query().Get("card")
	url1, flag1 := this.GetIntUrlParam("url1")
	url2, flag2 := this.GetIntUrlParam("url2", 100)
	url3, flag3 := this.GetIntPostParam("url3")
	url4, flag4 := this.GetIntPostParam("url4", 100)
	fmt.Fprintf(this.ResponseWriter, "\n%#v==%#v\n", url1, flag1)
	fmt.Fprintf(this.ResponseWriter, "\n%#v==%#v\n", url2, flag2)
	fmt.Fprintf(this.ResponseWriter, "\n%#v==%#v\n", url3, flag3)
	fmt.Fprintf(this.ResponseWriter, "\n%#v==%#v\n", url4, flag4)

}
