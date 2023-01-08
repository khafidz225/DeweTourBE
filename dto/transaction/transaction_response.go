package transactiondto

type TransactionResponse struct {
	ID         int    `json:"id" `
	Total      int    `json:"total" form:"total"`
	CounterQty int    `json:"counterqty" form:"counterqty"`
	Status     string `json:"status" form:"status"`
	TripID     int    `json:"trip_id" form:"trip_id"`
	// Image      string `json:"image" form:"image" grom:"type: varchar(255)"`
	UserID int `json:"user_id" form:"user_id"`
	// Trip       models.Trip `json:"trip"`
}
