package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type CardReceiver struct {
	Id         uint64 `orm:"fk;auto"`
	Card       string `orm:"size(50)"`
	Guid       string `orm:"size(50)"`
	Type       int8
	SchoolGuid string `orm:"size(50)"`
}
