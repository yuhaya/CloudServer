package models

import "github.com/astaxie/beego/orm"

type Classes struct {
	Id         uint64 `orm:"fk;auto"`
	Guid       string `orm:"unique;size(50)"`
	Name       string `orm:"size(20)"`
	SchoolGuid string `orm:"size(50)"`
}
