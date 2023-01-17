package handlers

import (
	dto "deweTourBE/dto/result"
	tripdto "deweTourBE/dto/trip"
	"deweTourBE/models"
	"deweTourBE/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

//var path_file = "http://localhost:5000/uploads/"

type handlerTrip struct {
	TripRepository repositories.TripRepository
}
//

// Untuk menampung Function menjadi satu
func HandlerTrip(TripRepository repositories.TripRepository) *handlerTrip {
	return &handlerTrip{TripRepository}
}

func (h *handlerTrip) FindTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	trip, err := h.TripRepository.FindTrip()
	// Image
	for i, p := range trip {
		trip[i].Image = path_file + p.Image
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//Encoder seperti kode morse di pramuka
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trip}
	//Encoder seperti kode morse di pramuka
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrip) GetTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	//untuk mengambil sesuai dengan idnya
	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	trip.Image = path_file + trip.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trip}
	//encoder seperti kode morse di pramuka
	json.NewEncoder(w).Encode(response)

}

func (h *handlerTrip) CreateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	fmt.Println(`ini user id`, userId)
	fmt.Println("ini user info", userInfo["role"])

	if userInfo["role"] != "admin" {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "besides admin can't create"}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)
	// filename = `http://localhost:5000/api/v1/` + filename

	country_id, _ := strconv.Atoi(r.FormValue("country_id"))
	Day, _ := strconv.Atoi(r.FormValue("day"))
	Night, _ := strconv.Atoi(r.FormValue("night"))
	Price, _ := strconv.Atoi(r.FormValue("price"))
	Quota, _ := strconv.Atoi(r.FormValue("quota"))
	request := tripdto.CreateTripRequest{
		Title:          r.FormValue("title"),
		CountryID:      country_id,
		Accomodation:   r.FormValue("accomodation"),
		Transportation: r.FormValue("transportation"),
		Eat:            r.FormValue("eat"),
		Day:            Day,
		Night:          Night,
		DateTrip:       r.FormValue("datetrip"),
		Price:          Price,
		Quota:          Quota,
		Description:    r.FormValue("description"),
		Image:          filename,
	}
	// if err := json.NewDecoder(r.).Decode(&request); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	//dari gthub validator
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	var ctx = context.Background()
	var CLOUD_NAME = "dfxarsquq"
	var API_KEY = "424662388976554"
	var API_SECRET = "izwGO6NvRBu5pNVJoPyp2j1oNC4"

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "DeweTour"})
	fmt.Println(resp.SecureURL)

	//mengambil dari models trip
	trip := models.Trip{
		Title:          request.Title,
		CountryID:      request.CountryID,
		Country:        models.Country{},
		Accomodation:   request.Accomodation,
		Transportation: request.Transportation,
		Eat:            request.Eat,
		Day:            request.Day,
		Night:          request.Night,
		Price:          request.Price,
		Quota:          request.Quota,
		Description:    request.Description,
		Image:          resp.SecureURL,
	}

	data, err := h.TripRepository.CreateTrip(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	datas, err := h.TripRepository.GetTrip(data.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: datas}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrip) UpdateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Countent-type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	request := tripdto.UpdateTripRequest{

		Image: filename,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	trip, err := h.TripRepository.GetTrip(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.FormValue("title") != "" {
		trip.Title = r.FormValue("title")
	}
	country_id, _ := strconv.Atoi(r.FormValue("country_id"))
	if r.FormValue("country_id") != "" {
		trip.CountryID = country_id
	}
	if r.FormValue("accomodation") != "" {
		trip.Accomodation = r.FormValue("accomodation")
	}
	if r.FormValue("transportation") != "" {
		trip.Transportation = r.FormValue("transportation")
	}
	if r.FormValue("eat") != "" {
		trip.Eat = r.FormValue("eat")
	}
	Day, _ := strconv.Atoi(r.FormValue("day"))
	if r.FormValue("day") != "" {
		trip.Day = Day
	}
	Night, _ := strconv.Atoi(r.FormValue("night"))
	if r.FormValue("night") != "" {
		trip.Night = Night
	}
	// if request.DateTrip.Format("2 January 2006") != "" {
	// 	trip.DateTrip = request.DateTrip
	// }
	Price, _ := strconv.Atoi(r.FormValue("price"))
	if r.FormValue("price") != "" {
		trip.Price = Price
	}
	Quota, _ := strconv.Atoi(r.FormValue("quota"))
	if r.FormValue("quota") != "" {
		trip.Quota = Quota
	}
	if r.FormValue("description") != "" {
		trip.Description = r.FormValue("description")
	}

	// dataContex := r.Context().Value("dataFile")
	// filename := dataContex.(string)
	if request.Image != "" {
		trip.Image = request.Image
	}

	data, err := h.TripRepository.UpdateTrip(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	data1, err := h.TripRepository.GetTrip(data.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data1}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerTrip) DeleteTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TripRepository.DeleteTrip(trip)
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

func tripConvertResponse(u models.Trip) tripdto.TripResponse {
	return tripdto.TripResponse{
		Title:          u.Title,
		CountryID:      u.CountryID,
		Country:        models.Country{},
		Accomodation:   u.Accomodation,
		Transportation: u.Transportation,
		Eat:            u.Eat,
		Day:            u.Day,
		Night:          u.Night,
		Price:          u.Price,
		Quota:          u.Quota,
		Description:    u.Description,
	}
}
