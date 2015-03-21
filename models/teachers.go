package models

type Teachers struct {
	Id         uint64 `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	Guid       string `sql:"unique;type:varchar(50);not null"`
	Name       string `sql:"type:varchar(50);not null"`
	Phone      string `sql:"type:varchar(20);not null"`
	Gender     int8   `sql:"default:0;not null"`
	SchoolGuid string `sql:"type:varchar(50);not null"`
}

func (this Teachers) TableName() string {
	return "ittr_teachers"
}
