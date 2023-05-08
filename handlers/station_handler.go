package handlers

import (
	"net/http"
	dto "server/dto/result"
	stationdto "server/dto/station"
	"server/models"
	"server/repositories"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerStation struct {
	StationRepository repositories.StationRepository
}

func HandlerStation(StationRepository repositories.StationRepository) *handlerStation {
	return &handlerStation{StationRepository}
}

func (h *handlerStation) FindAllStations(c echo.Context) error {
	stations, err := h.StationRepository.FindAllStations()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: stations})
}

func (h *handlerStation) GetStation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	station, err := h.StationRepository.GetStation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: station})
}

func (h *handlerStation) CreateStation(c echo.Context) error {
	request := new(stationdto.StationRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	station := models.Station{
		Name: request.Name,
		City: request.City,
	}

	newStation, err := h.StationRepository.CreateStation(station)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseStation(newStation)})
}

func (h *handlerStation) UpdateStation(c echo.Context) error {
	request := new(stationdto.StationRequest)

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	station, err := h.StationRepository.GetStation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		station.Name = request.Name
	}
	if request.City != "" {
		station.City = request.City
	}

	updatedStation, err := h.StationRepository.UpdateStation(station)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseStation(updatedStation)})
}

func (h *handlerStation) DeleteStation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	station, err := h.StationRepository.GetStation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	deletedStation, err := h.StationRepository.DeleteStation(station)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: deletedStation})
}
func convertResponseStation(u models.Station) models.Station {
	return models.Station{
		ID:   u.ID,
		Name: u.Name,
		City: u.City,
	}
}
