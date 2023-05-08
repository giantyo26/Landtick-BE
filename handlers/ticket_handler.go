package handlers

import (
	"fmt"
	"net/http"
	dto "server/dto/result"
	ticketdto "server/dto/ticket"
	"server/models"
	"server/repositories"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerTicket struct {
	TicketRepository repositories.TicketRepository
}

func HandlerTicket(TicketRepository repositories.TicketRepository) *handlerTicket {
	return &handlerTicket{TicketRepository}
}

func (h *handlerTicket) FindAllTickets(c echo.Context) error {
	tickets, err := h.TicketRepository.FindAllTickets()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: tickets})
}

func (h *handlerTicket) GetTicket(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ticket, err := h.TicketRepository.GetTicket(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTicket(ticket)})
}

func (h *handlerTicket) CreateTicket(c echo.Context) error {
	startStationId, _ := strconv.Atoi(c.FormValue("start_station_id"))
	destinationStationId, _ := strconv.Atoi(c.FormValue("destination_station_id"))
	departureDate, _ := time.Parse("2006-01-02", c.FormValue("departure_date"))
	startTime := c.FormValue("start_time")
	arrivalTime := c.FormValue("arrival_time")
	price, _ := strconv.Atoi(c.FormValue("price"))
	qty, _ := strconv.Atoi(c.FormValue("qty"))

	request := ticketdto.TicketRequest{
		TrainName:            c.FormValue("train_name"),
		TrainClass:           c.FormValue("train_class"),
		StartStationId:       startStationId,
		DestinationStationId: destinationStationId,
		DepartureDate:        departureDate,
		StartTime:            startTime,
		ArrivalTime:          arrivalTime,
		Price:                price,
		Qty:                  qty,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	startTimeObj, _ := time.Parse("15:04", startTime)
	ArrivalTimeObj, _ := time.Parse("15:04", arrivalTime)

	trainDuration := CountTrainDuration(startTimeObj, ArrivalTimeObj)

	ticket := models.Ticket{
		TrainName:            request.TrainName,
		TrainClass:           request.TrainClass,
		StartStationId:       request.StartStationId,
		DestinationStationId: request.DestinationStationId,
		DepartureDate:        request.DepartureDate,
		StartTime:            request.StartTime,
		ArrivalTime:          request.ArrivalTime,
		TrainDuration:        trainDuration,
		Price:                request.Price,
		Qty:                  request.Qty,
	}

	addedTicket, err := h.TicketRepository.CreateTicket(ticket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: addedTicket})
}

func (h *handlerTicket) DeleteTicket(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param(("id")))

	ticket, err := h.TicketRepository.GetTicket(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	deletedTicket, err := h.TicketRepository.DeleteTicket(ticket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: deletedTicket})
}

func (h *handlerTicket) FilterTickets(c echo.Context) error {
	// Get the query parameters
	departureDateStr := c.QueryParam("departure_date")
	startStationIDParam := c.QueryParam("start_station_id")
	destinationStationIDParam := c.QueryParam("destination_station_id")

	// Parse the departure date 
	var departureDate time.Time
	if departureDateStr != "" {
		var err error
		departureDate, err = time.Parse("2006-01-02", departureDateStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Invalid date format. Use format yyyy-mm-dd"})
		}
	}

	// Convert the station id to int
	startStationID, err := strconv.Atoi(startStationIDParam)
	if err != nil && startStationIDParam != "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Invalid start_station_id"})
	}

	destinationStationID, err := strconv.Atoi(destinationStationIDParam)
	if err != nil && destinationStationIDParam != "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Invalid destination_station_id"})
	}

	// Call the FindAllTickets method of the ticket repository
	tickets, err := h.TicketRepository.FilterTickets(departureDate, startStationID, destinationStationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: tickets})
}

func CountTrainDuration(StartTime, ArrivalTime time.Time) string {
	trainDuration := ArrivalTime.Sub(StartTime)
	durationHours := int(trainDuration.Hours())
	durationMinutes := int(trainDuration.Minutes()) % 60

	trainDurationString := fmt.Sprintf("%d hours %d minutes", durationHours, durationMinutes)

	return trainDurationString
}

func convertResponseTicket(u models.Ticket) ticketdto.TicketResponse {
	return ticketdto.TicketResponse{
		ID:                   u.ID,
		TrainName:            u.TrainName,
		TrainClass:           u.TrainClass,
		StartStationId:       u.StartStationId,
		StartStation:         u.StartStation,
		DestinationStationId: u.DestinationStationId,
		DestinationStation:   u.DestinationStation,
		DepartureDate:        u.DepartureDate,
		StartTime:            u.StartTime,
		ArrivalTime:          u.ArrivalTime,
		TrainDuration:        u.TrainDuration,
		Price:                u.Price,
		Qty:                  u.Qty,
	}
}
