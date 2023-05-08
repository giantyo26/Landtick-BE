package repositories

import (
	"server/models"
	"time"

	"gorm.io/gorm"
)

type TicketRepository interface {
	FindAllTickets() ([]models.Ticket, error)
	GetTicket(ID int) (models.Ticket, error)
	CreateTicket(ticket models.Ticket) (models.Ticket, error)
	DeleteTicket(ticket models.Ticket) (models.Ticket, error)
	FilterTickets(date time.Time, startStation int, destinationStation int) ([]models.Ticket, error)
}

func RepositoryTicket(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllTickets() ([]models.Ticket, error) {
	var tickets []models.Ticket
	err := r.db.Preload("StartStation").Preload("DestinationStation").Find(&tickets).Error

	return tickets, err
}

func (r *repository) GetTicket(ID int) (models.Ticket, error) {
	var ticket models.Ticket
	err := r.db.Preload("StartStation").Preload("DestinationStation").First(&ticket, "id = ?", ID).Error

	return ticket, err
}

func (r *repository) CreateTicket(ticket models.Ticket) (models.Ticket, error) {
	err := r.db.Preload("StartStation").Preload("DestinationStation").Create(&ticket).Error
	return ticket, err
}

func (r *repository) DeleteTicket(ticket models.Ticket) (models.Ticket, error) {
	err := r.db.Delete(&ticket).Error

	return ticket, err
}

func (r *repository) FilterTickets(departureDate time.Time, startStationId int, destinationStationId int) ([]models.Ticket, error) {
	var tickets []models.Ticket
	query := r.db.Preload("StartStation").Preload("DestinationStation")

	if !departureDate.IsZero() {
		query = query.Where("departure_date = ?", departureDate)
	}

	if startStationId != 0 {
		query = query.Where("start_station_id = ?", startStationId)
	}

	if destinationStationId != 0 {
		query = query.Where("destination_station_id = ?", destinationStationId)
	}

	err := query.Find(&tickets).Error

	return tickets, err
}
