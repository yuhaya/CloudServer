package models

type ClassTeacher struct {
	Id          uint64 `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	ClassGuid   string `sql:"type:varchar(50);not null"`
	TeacherGuid string `sql:"type:varchar(50);not null"`
	Adviser     int8   `sql:"default:0;not null"`
	SchoolGuid  string `sql:"type:varchar(50);not null"`
}

func (this ClassTeacher) TableName() string {
	return "ittr_class_teacher"
}
