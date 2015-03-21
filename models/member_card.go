package models

type MemberCard struct {
	Id         uint64 `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	Card       string `orm:"type:varchar(50)"`
	Guid       string `orm:"type:varchar(50)"`
	Type       int8   `orm:"default:1"`
	SchoolGuid string `orm:"type:varchar(50)"`
}

func (this MemberCard) TableName() string {
	return "ittr_member_card"
}
