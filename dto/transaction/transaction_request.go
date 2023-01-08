package transactiondto

type CreateTransactionRequest struct {
	ID         int    `json:"id" `
	Total      int    `json:"total" form:"total"`
	CounterQty int    `json:"counterqty" form:"counterqty"`
	Status     string `json:"status" form:"status"`
	TripID     int    `json:"trip_id" form:"trip_id"`
	UserID     int    `json:"user_id" form:"user_id"`
	// Image      string `json:"image" form:"image" grom:"type: varchar(255)"`
	// Trip       models.Trip `json:"trip"`
}

type UpdateTransactionRequest struct {
	ID         int    `json:"id" `
	Total      int    `json:"total" form:"total"`
	CounterQty int    `json:"counterqty" form:"counterqty"`
	Status     string `json:"status" form:"status"`
	TripID     int    `json:"trip_id" form:"trip_id"`
	UserID     int    `json:"user_id" form:"user_id"`
	// Image      string `json:"image" form:"image" grom:"type: varchar(255)"`
	// Trip       models.Trip `json:"trip"`
}
