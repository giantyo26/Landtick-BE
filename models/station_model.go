package models

type Station struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" gorm:"type: varchar(100)"`
	City string `json:"city" gorm:"type: varchar(255)"`
}
