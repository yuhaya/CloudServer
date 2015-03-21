package models

import (
	"time"
)

type Admins struct {
	Id         int       `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	Guid       string    `sql:"unique;type:varchar(50);not null"`
	Username   string    `sql:"type:varchar(50);not null"`
	Password   string    `sql:"type:varchar(50);not null"`
	Realname   string    `sql:"type:varchar(50);null"`
	SchoolGuid string    `sql:"type:varchar(50)"`
	CreateTime time.Time `sql:"type:datetime;not null"`
	Super      int8      `sql:"default:0;not null"`
	Enabled    int8      `sql:"default:1;not null"`
}

func (this Admins) TableName() string {
	return "ittr_admins"
}
