package userdto

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	NIK      string `json:"nik" gorm:"type: varchar(255)"`
}
