package models

import (
	"time"
)

type Users struct {
	Id         uint64    `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	Guid       string    `sql:"unique;type:varchar(50);not null"`
	Phone      string    `sql:"type:varchar(20);null"`
	Realname   string    `sql:"type:varchar(10);null"`
	Password   string    `sql:"type:varchar(50);not null"`
	Spell      string    `sql:"type:varchar(10);null"`
	Gender     int8      `sql:"default:1;not null"`
	IdCard     string    `sql:"type:varchar(20);null"`
	Picture    string    `sql:"type:varchar(100);null"`
	SchoolGuid string    `sql:"type:varchar(50);not null"`
	CreateTime time.Time `sql:"type:datetime;not null"`
}

func (this Users) TableName() string {
	return "ittr_users"
}
