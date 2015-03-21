package models

type Families struct {
	Id         uint64 `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	Name       string `sql:"type:varchar(50);not null"`
	Guid       string `sql:"unique;type:varchar(50);not null"`
	FirstGuid  string `sql:"type:varchar(50);not null"`
	SchoolGuid string `sql:"type:varchar(50);not null"`
}

func (this Families) TableName() string {
	return "ittr_families"
}
