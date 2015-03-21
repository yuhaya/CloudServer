package models

type Grades struct {
	Id         uint64 `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	Guid       string `sql:"unique;type:varchar(50);not null"`
	Name       string `sql:"type:varchar(50);not null"`
	Rating     uint64 `sql:"default:0;not null"`
	SchoolGuid string `sql:"type:varchar(50);not null"`
}

func (this Grades) TableName() string {
	return "ittr_grades"
}
