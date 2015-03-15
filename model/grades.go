package models

import "github.com/astaxie/beego/orm"

type Grades struct {
	Id         uint64 `orm:"fk;auto"`
	Guid       string `orm:"unique;size(50)"`
	Name       string `orm:"size(50)"`
	Rating     uint64 `orm:"default(0)"`
	SchoolGuid string `orm:"size(50)"`
}
