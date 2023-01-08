package models

type Trip struct {
	ID             int     `json:"id" gorm:"primary_key:auto_increment"`
	Title          string  `json:"title" gorm:"type: varchar(255)"`
	CountryID      int     `json:"country_id" form:"country_id"`
	Country        Country `json:"country" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Accomodation   string  `json:"accomodation" gorm:"type: varchar(255)"`
	Transportation string  `json:"transportation" gorm:"type: varchar(255)"`
	Eat            string  `json:"eat" gorm:"type: varchar(255)"`
	Day            int     `json:"day" gorm:"type: int"`
	Night          int     `json:"night" gorm:"type: int"`
	DateTrip       string  `json:"datetrip"`
	Price          int     `json:"price" gorm:"type: int"`
	Quota          int     `json:"quota" gorm:"type: int"`
	Description    string  `json:"description" form:"description" gorm:"type: varchar(255)"`
	Image          string  `json:"image" form:"image" grom:"type: varchar(255)"`
}

func (Trip) TableName() string {
	return "trip"
}
