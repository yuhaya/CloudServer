package models

import (
	//	"strings"
	"time"
)

type Admins struct {
	Id         int64  `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Guid       string `sql:"unique;type:varchar(50)"`
	Username   string `sql:"type:varchar(50)"`
	Password   string `sql:"type:varchar(50)"`
	Realname   string `sql:"type:varchar(50);null"`
	SchoolGuid string `sql:"type:varchar(50)"`
	CreateTime time.Time
	Super      int8 `sql:"default:0"`
	Enabled    int8 `sql:"default:1"`
}
