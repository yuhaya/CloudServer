package models

type GradeClass struct {
	Id         uint64 `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	GradeGuid  string `sql:"type:varchar(50);not null"`
	ClassGuid  string `sql:"type:varchar(50);not null"`
	SchoolGuid string `sql:"type:varchar(50);not null"`
}

func (this GradeClass) TableName() string {
	return "ittr_grade_class"
}
