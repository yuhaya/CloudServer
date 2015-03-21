package models

type System struct {
	Id         uint64 `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	Key        string `sql:"type:varchar(50);not null"`
	Value      string `sql:"size:250;not null"`
	SchoolGuid string `sql:"type:varchar(50);not null"`
}

func (this System) TableName() string {
	return "ittr_system"
}
