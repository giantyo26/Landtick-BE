package ticketdto

import "time"

type TicketRequest struct {
	TrainName            string    `json:"train_name" form:"train_name" validate:"required"`
	TrainClass           string    `json:"train_class" form:"train_class" validate:"required"`
	StartStationId       int       `json:"start_station_id" form:"start_station_id" validate:"required"`
	DestinationStationId int       `json:"destination_station_id" form:"destination_station_id" validate:"required"`
	DepartureDate        time.Time `json:"departure_date" form:"departure_date" validate:"required"`
	StartTime            string    `json:"start_time" form:"start_time" validate:"required"`
	ArrivalTime          string    `json:"arrival_time" form:"arrival_time" validate:"required"`
	Price                int       `json:"price" form:"price" validate:"required"`
	Qty                  int       `json:"qty" form:"qty" validate:"required"`	
}
