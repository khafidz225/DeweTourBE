package countrydto

type CreateCountryRequest struct {
	//json hasil response di postman
	//form untuk menerima inputan dari mananya
	//validate perintah harus di isi jika tidak di isi akan eror
	Name string `json:"name" form:"name" validate:"required"`
}

type UpdateCountryRequest struct {
	Name string `json:"name" form:"name"`
}
