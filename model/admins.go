package models

import (
	"github.com/astaxie/beego/validation"
	"strings"
	"time"
)

type Admins struct {
	Id         uint64    `orm:"fk;auto"`
	Guid       string    `orm:"unique;size(50)" valid:"Required"`
	Username   string    `orm:"size(50)" valid:"Required"`
	Password   string    `orm:"size(50)" valid:"Required"`
	Realname   string    `orm:"size(50);null"`
	SchoolGuid string    `orm:"size(50)" valid:"Required"`
	CreateTime time.Time `orm:"type(datetime)" valid:"Required"`
	Super      int8      `orm:"default(0)" valid:"Required;Range(0, 1)"`
	Enabled    int8      `orm:"default(1)" valid:"Required;Range(0, 1)""`
}
