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

}
