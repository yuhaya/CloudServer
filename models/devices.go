package models

import (
//	"LocalServer/tool"
//	"fmt"
//	"github.com/astaxie/beego/orm"
)

type Devices struct {
	Id          uint64 `orm:"fk;auto"`
	Guid        string `orm:"unique;size(50)"`
	Device      string `orm:"size(50)"`
	Kind        int8   `orm:"default(0)"`
	Vmp         string `orm:"size(10)"`
	SchoolGuid  string `orm:"size(50)"`
	Group       int8   `orm:"default(0)"`
	Description string `orm:"size(255)"`
	Status      int8   `orm:"default(1)"`
	Enabled     int8   `orm:"default(1)"`
}
