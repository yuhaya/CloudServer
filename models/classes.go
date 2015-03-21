package models

type Classes struct {
	Id         uint64 `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	Guid       string `sql:"unique;type:varchar(50);not null"`
	Name       string `sql:"type:varchar(20);not null"`
	SchoolGuid string `sql:"type:varchar(50);not null"`
}

func (this Classes) TableName() string {
	return "ittr_classes"
}
