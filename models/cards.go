package models

type Cards struct {
	Id         uint64 `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	Guid       string `sql:"unique;type:varchar(50);not null"`
	Kind       int8   `sql:"default:1;not null"`
	SchoolGuid string `sql:"type:varchar(50);not null"`
	FamilyGuid string `sql:"type:varchar(50);not null"`
	Enabled    int8   `sql:"default:1;not null"`
}

func (this Cards) TableName() string {
	return "ittr_cards"
}
