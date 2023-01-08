package tripdto

import (
	"deweTourBE/models"
)

type TripResponse struct {
	Title          string         `json:"title" form:"title" grom:"type: varchar(255)"`
	CountryID      int            `json:"country_id" form:"country_id"`
	Country        models.Country `json:"country" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Accomodation   string         `json:"accomodation" form:"accomodation" grom:"type: varchar(255)"`
	Transportation string         `json:"transportation" form:"transportation" grom:"type: varchar(255)"`
	Eat            string         `json:"eat" form:"eat" grom:"type: varchar(255)"`
	Day            int            `json:"day" form:"day" grom:"type: int"`
	Night          int            `json:"night" form:"night" grom:"type: int"`
	DateTrip       string         `json:"datetrip" form:"datetrip"`
	Price          int            `json:"price" form:"price" grom:"type: int"`
	Quota          int            `json:"quota" form:"quota" grom:"type: int"`
	Description    string         `json:"description" form:"description" grom:"type: varchar(255)"`
	Image          string         `json:"image" form:"image" grom:"type: varchar(255)"`
}
