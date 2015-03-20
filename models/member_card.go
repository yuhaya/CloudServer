package models

import (
//	"fmt"
//	"github.com/astaxie/beego/orm"
)

type MemberCard struct {
	Id         uint64 `orm:"fk;auto"`
	Card       string `orm:"size(50)"`
	Guid       string `orm:"size(50)"`
	Type       int8   `orm:"default(1)"`
	SchoolGuid string `orm:"size(50)"`
}
