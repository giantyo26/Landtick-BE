package ticketdto

import (
	"server/models"
	"time"
)

type TicketResponse struct {
	ID                   int            `json:"id"`
	TrainName            string         `json:"train_name"`
	TrainClass           string         `json:"train_class"`
	StartStationId       int            `json:"start_station_id"`
	StartStation         models.Station `json:"start_station"`
	DestinationStationId int            `json:"destination_station_id"`
	DestinationStation   models.Station `json:"destination_station"`
	DepartureDate        time.Time      `json:"departure_date"`
	StartTime            string         `json:"start_time"`
	ArrivalTime          string         `json:"arrival_time"`
	Price                int            `json:"price"`
	TrainDuration        string         `json:"train_duration"`
	Qty                  int            `json:"qty"`
}
