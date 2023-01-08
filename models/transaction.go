package models

type Transaction struct {
	ID         int    `json:"id" gorm:"primary_key:auto_increment"`
	Total      int    `json:"total" form:"total"`
	CounterQty int    `json:"counterqty" form:"counterqty"`
	Status     string `json:"status" form:"status"`
	TripID     int    `json:"trip_id" form:"trip_id"`
	// Image      string        `json:"image" form:"image" grom:"type: varchar(255)"`
	Trip   Trip          `json:"trip" form:"trip"`
	UserID int           `json:"user_id" form:"user_id"`
	User   UsersResponse `json:"user"`
}

func (Transaction) TableName() string {
	return "transactions"
}
