package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	FullName string `json:"fullname" grom:"type: varchar(255)"`
	Email    string `json:"email" grom:"type: varchar(255)"`
	Password string `json:"password" grom:"type: varchar(255)"`
	Phone    string `json:"phone" grom:"type: varchar(255)"`
	Address  string `json:"address" grom:"type: varchar(255)"`
	Role     string `json:"role" `
}

type UsersResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}

func (UsersResponse) TableName() string {
	return "users"
}
