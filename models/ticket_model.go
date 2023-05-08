package models

import "time"

type Ticket struct {
	ID                   int       `json:"id" gorm:"primary_key:auto_increment"`
	TrainName            string    `json:"train_name" form:"train_name"`
	TrainClass           string    `json:"train_class" form:"train_class"`
	StartStationId       int       `json:"start_station_id" form:"start_station_id"`
	StartStation         Station   `json:"start_station" form:"start_station" gorm:"foreignKey:StartStationId"`
	DestinationStationId int       `json:"destination_station_id" form:"destination_station_id" gorm:"foreignKey:DestinationStationId"`
	DestinationStation   Station   `json:"destination_station" form:"destination_station"`
	DepartureDate        time.Time `json:"departure_date" form:"departure_date"`
	StartTime            string    `json:"start_time" form:"start_time"`
	ArrivalTime          string    `json:"arrival_time" form:"arrival_time"`
	TrainDuration        string    `json:"train_duration"`
	Price                int       `json:"price" form:"price"`
	Qty                  int       `json:"qty" form:"qty"`
}
