package models

import (
	"time"
)

type Attendances struct {
	Id         uint64    `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Card       string    `sql:"type:varchar(50);not null"`
	Time       time.Time `sql:"type:datetime;not null"`
	Type       int8      `sql:"default:0;not null"`
	SchoolGuid string    `sql:"type:varchar(50);not null"`
	Auto       int8      `sql:"default:1;not null"`
}

func (this Attendances) TableName() string {
	return "ittr_attendances"
}
