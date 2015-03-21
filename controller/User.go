package controller

import (
	"fmt"
)

type User struct {
	Base
}

func (this User) ActionHello() {
	fmt.Fprint(this.ResponseWriter, "Hello Go !")
}
