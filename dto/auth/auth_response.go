package authdto

type RegisterResponse struct {
	Fullname string `json:"fullname"  gorm:"type: varchar(255)"`
	Username string `json:"username" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Token    string `json:"token" gorm:"type: varchar(255)"`
}

type LoginResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"  gorm:"type: varchar(255)"`
	Username string `json:"username" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Token    string `json:"token" gorm:"type: varchar(255)"`
}
