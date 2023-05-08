package database

import (
	"fmt"
	"server/models"
	"server/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{}, &models.Ticket{}, &models.Station{}, &models.Transaction{}, models.City{})
	if err != nil {
		fmt.Println(err)
		panic("Migration failed")
	}
	fmt.Println("Migration sucess")
}
