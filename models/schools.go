package models

import (
	"time"
)

type Schools struct {
	Id         uint64    `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	Guid       string    `sql:"unique;type:varchar(50)"`
	Name       string    `sql:"type:varchar(30)"`
	Spell      string    `sql:"type:varchar(50)"`
	Province   string    `sql:"type:varchar(20)"`
	City       string    `sql:"type:varchar(20)"`
	County     string    `sql:"type:varchar(20)"`
	Address    string    `sql:"type:varchar(80);null"`
	UpdateTime time.Time `sql:"type:datetime"`
	DoorNum    int8      `sql:"default:1"`
	Enabled    int8      `sql:"default:1"`
}

func (this Schools) TableName() string {
	return "ittr_schools"
}
