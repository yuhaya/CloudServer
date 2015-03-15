package models

import (
	"LocalServer/tool"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

type Families struct {
	Id         uint64 `orm:"fk;auto"`
	Name       string `orm:"size(50)"`
	Guid       string `orm:"unique;size(50)"`
	FirstGuid  string `orm:"size(50)"`
	SchoolGuid string `orm:"size(50)"`
}
