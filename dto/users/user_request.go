package usersdto

type CreateUserRequest struct {
	//json hasil response di postman
	//form untuk menerima inputan dari mananya
	//validate perintah harus di isi jika tidak di isi akan eror
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Role     string `json:"role" form:"role"`
}

type UpdateUserRequest struct {
	FullName string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone" `
	Address  string `json:"address" form:"address" `
	Role     string `json:"role" form:"role"`
}
