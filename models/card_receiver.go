package models

type CardReceiver struct {
	Id         uint64 `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	Card       string `sql:"type:varchar(50);not null"`
	Guid       string `sql:"type:varchar(50);not null"`
	Type       int8   `sql:"not null"`
	SchoolGuid string `sql:"type:varchar(50);not null"`
}

func (this CardReceiver) TableName() string {
	return "ittr_card_receiver"
}
