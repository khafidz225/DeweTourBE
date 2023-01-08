package models

type Country struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}

func (Country) TableName() string {
	return "countrys"
}
