package countrydto

type CountryResponse struct {
	Name string `json:"name" form:"name" validate:"required"`
}
