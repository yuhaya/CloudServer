package models

type FamilyRelation struct {
	Id         uint64 `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	UserGuid   string `sql:"type:varchar(50);not null"`
	StuGuid    string `sql:"type:varchar(50);not null"`
	Relation   string `sql:"type:varchar(50);not null"`
	SchoolGuid string `sql:"type:varchar(50);not null"`
}

func (this FamilyRelation) TableName() string {
	return "ittr_family_relation"
}
