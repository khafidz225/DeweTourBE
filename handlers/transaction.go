package handlers

import (
	dto "deweTourBE/dto/result"
	transactiondto "deweTourBE/dto/transaction"
	"deweTourBE/models"
	"deweTourBE/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gopkg.in/gomail.v2"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

// Untuk menampung function menjadi satu
func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) FindTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	transaction, err := h.TransactionRepository.FindTransaction()
	// for i, p := range transaction {
	// 	transaction[i].Image = path_file + p.Image
	// }
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//Encoder seperti kode morse di pramuka
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	//Encoder seperti kode morse di pramuka
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	//untuk mengambil sesuai dengan idnya
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// transaction.Image = path_file + transaction.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	// ID         int
	// CounterQty int
	// Total      int
	// Status     string
	// Attachment string
	// TripID     int

	// dataContex := r.Context().Value("dataFile")
	// filename := dataContex.(string)

	// CounterQty, _ := strconv.Atoi(r.FormValue("counterqty"))
	// Total, _ := strconv.Atoi(r.FormValue("total"))
	// trip_id, _ := strconv.Atoi(r.FormValue("trip_id"))
	// user_id, _ := strconv.Atoi(r.FormValue("user_id"))

	// request := transactiondto.CreateTransactionRequest{
	// 	Total:      ,
	// 	CounterQty: CounterQty,
	// 	Status:     r.FormValue("status"),
	// 	UserID:     user_id,
	// 	// Image:      filename,
	// 	TripID: trip_id,
	// }

	request := new(transactiondto.CreateTransactionRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var transIdIsMatch = false
	var transactionId int

	for !transIdIsMatch {
		transactionId = int(time.Now().Unix())
		transaction, _ := h.TransactionRepository.GetTransaction(transactionId)
		if transaction.ID == 0 {
			transIdIsMatch = true
		}
	}

	transaction := models.Transaction{
		ID:         transactionId,
		Total:      request.Total,
		CounterQty: request.CounterQty,
		Status:     request.Status,
		// Image:      request.Image,
		UserID: request.UserID,
		User:   models.UsersResponse{},
		TripID: request.TripID,
		Trip:   models.Trip{},
	}
	fmt.Println(transaction)

	dataTransaction, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	transactionGet, err := h.TransactionRepository.GetTransaction(dataTransaction.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transactionGet.ID),
			GrossAmt: int64(transactionGet.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: transactionGet.User.FullName,
			Email: transactionGet.User.Email,
		},
	}
	snapResp, _ := s.CreateTransaction(req)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: snapResp}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) Notification(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	convertId, _ := strconv.Atoi(orderId)

	transaction, _ := h.TransactionRepository.GetTransaction(convertId)
	fmt.Println(transactionStatus, fraudStatus, orderId, transaction)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			SendEmail("pending", transaction)
			h.TransactionRepository.UpdateTransaction("pending", transaction)
		} else if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			SendEmail("success", transaction)
			h.TransactionRepository.UpdateTransaction("success", transaction)
		}
	} else if transactionStatus == "settlement" {
		SendEmail("success", transaction)
		h.TransactionRepository.UpdateTransaction("success", transaction)
	} else if transactionStatus == "deny" {
		SendEmail("failed", transaction)
		h.TransactionRepository.UpdateTransaction("failed", transaction)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		SendEmail("failed", transaction)
		h.TransactionRepository.UpdateTransaction("failed", transaction)
	} else if transactionStatus == "pending" {
		SendEmail("pending", transaction)
		h.TransactionRepository.UpdateTransaction("pending", transaction)
	}
	w.WriteHeader(http.StatusOK)
}

func SendEmail(status string, transaction models.Transaction) {
	var CONFIG_SMTP_HOST = "smtp.gmail.com"
	var CONFIG_SMTP_PORT = 587
	var CONFIG_SENDER_NAME = "DeweTour <khaaafidz225@gmail.com>"
	var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
	var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

	var tripTitle = transaction.Trip.Title
	var total = strconv.Itoa(transaction.Total)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", transaction.User.Email)
	mailer.SetHeader("Subject", "Status Transaksi ")
	mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
    <html lang="en">
      <head>
      <meta charset="UTF-8" />
      <meta http-equiv="X-UA-Compatible" content="IE=edge" />
      <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      <title>Document</title>
      <style>
        h1 {
        color: brown;
        }
      </style>
      </head>
      <body>
      <h2>Product payment :</h2>
      <ul style="list-style-type:none;">
        <li>Name : %s</li>
        <li>Total payment: Rp.%s</li>
        <li>Status : <b>%s</b></li>
        <li>Iklan : <b>%s</b></li>
      </ul>
      </body>
    </html>`, tripTitle, total, status, "khafidz merch potongan harga 50% ✌"))

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Pesan Sudah Terkirim Kawan")
}

// func (h *handlerTransaction) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")

// dataContexs := r.Context().Value("dataFile")
// filename := dataContexs.(string)

// request := transactiondto.UpdateTransactionRequest{
// 	Image: filename,
// }

// id, _ := strconv.Atoi(mux.Vars(r)["id"])
// transaction, err := h.TransactionRepository.GetTransaction(int(id))
// if err != nil {
// 	w.WriteHeader(http.StatusBadRequest)
// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 	json.NewEncoder(w).Encode(response)
// 	return
// }
// Total      int    `json:"total" form:"total"`
// CounterQty int    `json:"counterqty" form:"counterqty"`
// Status     string `json:"status" form:"status"`
// Image      string `json:"image" form:"image" grom:"type: varchar(255)"`
// TripID     int    `json:"trip_id" form:"trip_id"`
// Trip       Trip   `json:"trip" form:"trip"`

// Total, _ := strconv.Atoi(r.FormValue("total"))
// if r.FormValue("total") != "" {
// 	transaction.Total = Total
// }
// CounterQty, _ := strconv.Atoi(r.FormValue("counterqty"))
// if r.FormValue("counterqty") != "" {
// 	transaction.CounterQty = CounterQty
// }
// if r.FormValue("status") != "" {
// 	transaction.Status = r.FormValue("status")
// }
// if request.Image != "" {
// 	transaction.Image = request.Image
// }

// 	TripID, _ := strconv.Atoi(r.FormValue("trip_id"))
// 	if r.FormValue("trip_id") != "" {
// 		transaction.TripID = TripID
// 	}

// 	data, err := h.TransactionRepository.UpdateTransaction(transaction)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	datas, err := h.TransactionRepository.GetTransaction(data.ID)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: datas}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlerTransaction) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TransactionRepository.DeleteTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func transactionConvertResponse(u models.Transaction) transactiondto.TransactionResponse {
	return transactiondto.TransactionResponse{
		Total:      u.Total,
		CounterQty: u.CounterQty,
		Status:     u.Status,
		// Image:      u.Image,
		TripID: u.TripID,
	}
}
