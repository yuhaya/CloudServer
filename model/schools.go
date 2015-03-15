package models

import (
	"github.com/astaxie/beego/orm"
	"time"

	"fmt"
	"strconv"
)

type Schools struct {
	Id         uint64    `orm:"fk;auto"`
	Guid       string    `orm:"unique;size(50)"`
	Name       string    `orm:"size(30)"`
	Spell      string    `orm:"size(50)"`
	Province   string    `orm:"size(20)"`
	City       string    `orm:"size(20)"`
	County     string    `orm:"size(20)"`
	Address    string    `orm:"size(80);null"`
	UpdateTime time.Time `orm:"type(datetime)"`
	DoorNum    int8      `orm:"default(1)"`
	Enabled    int8      `orm:"default(1)"`
}