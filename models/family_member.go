package models

type FamilyMember struct {
	Id         uint64 `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	FamilyGuid string `sql:"type:varchar(50);not null"`
	MemberGuid string `sql:"type:varchar(50);not null"`
	Type       int8   `sql:"default:1;not null"`
	SchoolGuid string `sql:"type:varchar(50);not null"`
}

func (this FamilyMember) TableName() string {
	return "ittr_family_member"
}
