package models

//import "github.com/astaxie/beego/orm"

type Teachers struct {
	Id         uint64 `orm:"fk;auto"`
	Guid       string `orm:"unique;size(50)"`
	Name       string `orm:"size(50)"`
	Phone      string `orm:"size(20)"`
	Gender     int8   `orm:"default(0)"`
	SchoolGuid string `orm:"size(50)"`
}
