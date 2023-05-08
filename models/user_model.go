package models

type User struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" gorm:"type: varchar(255)"`
	Username string `json:"username" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	NIK      string `json:"nik" gorm:"type: varchar(255)"`
	Role     string `json:"role" gorm:"default:user;type: varchar(20)"`
}
