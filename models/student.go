package models

import (
	"time"
)

type Students struct {
	Id          uint64    `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	Guid        string    `sql:"unique;type:varchar(50)"`
	Sid         string    `sql:"type:varchar(50);null"`
	Realname    string    `sql:"type:varchar(10)"`
	Spell       string    `sql:"type:varchar(10)"`
	Gender      int8      `sql:"default:1"`
	Picture     string    `sql:"size:255;null"`
	Birthday    time.Time `sql:"type:date"`
	School_guid string    `sql:"type:varchar(50);null"`
	Grade_guid  string    `sql:"type:varchar(50);null"`
	Class_guid  string    `sql:"type:varchar(20);null"`
	Enrol_time  time.Time `sql:"type:date"`
	Create_time time.Time `sql:"type:datetime"`
}

func (this Students) TableName() string {
	return "ittr_students"
}
