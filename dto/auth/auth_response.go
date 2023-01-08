package authdto

type LoginResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Token    string `json:"token" gorm:"type: varchar(255)"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type CheckAuthResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Role     string `json:"role" gorm:"type:varchar(255)"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}
