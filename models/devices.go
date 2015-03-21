package models

type Devices struct {
	Id          uint64 `sql:"AUTO_INCREMENT;not null" gorm:"primary_key"`
	Guid        string `sql:"unique;type:varchar(50);not null"`
	Device      string `sql:"type:varchar(50);not null"`
	Kind        int8   `sql:"default:0;not null"`
	Vmp         string `sql:"type:varchar(10);not null"`
	SchoolGuid  string `sql:"type:varchar(50);not null"`
	Group       int8   `sql:"default(0);not null"`
	Description string `sql:"size(255);not null"`
	Status      int8   `sql:"default:1;not null"`
	Enabled     int8   `sql:"default:1;not null"`
}

func (this Devices) TableName() string {
	return "ittr_devices"
}
