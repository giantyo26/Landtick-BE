package routes

import (
	"server/handlers"
	"server/pkg/mysql"
	"server/repositories"

	"github.com/labstack/echo/v4"
)

func StationRoute(e *echo.Group) {
	StationRepository := repositories.RepositoryStation(mysql.DB)

	h := handlers.HandlerStation(StationRepository)

	e.GET("/stations", h.FindAllStations)
	e.GET("/stations/:id", h.GetStation)
	e.POST("/stations", h.CreateStation)
	e.PATCH("/stations/:id", h.UpdateStation)
	e.DELETE("/stations/:id", h.DeleteStation)
}
