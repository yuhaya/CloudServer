package models

import (
	"github.com/astaxie/beego/orm"
)

type Cards struct {
	Id         uint64 `orm:"fk;auto"`
	Guid       string `orm:"unique;size(50)"`
	Kind       int8   `orm:"default(1)"`
	SchoolGuid string `orm:"size(50)"`
	FamilyGuid string `orm:"size(50)"`
	Enabled    int8   `orm:"default(1)"`
}
