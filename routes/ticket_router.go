package routes

import (
	"server/handlers"
	"server/pkg/mysql"
	"server/repositories"

	"github.com/labstack/echo/v4"
)

func TicketRoute(e *echo.Group) {
	TicketRepository := repositories.RepositoryTicket(mysql.DB)

	h := handlers.HandlerTicket(TicketRepository)

	e.GET("/tickets", h.FindAllTickets)
	e.GET("/tickets/:id", h.GetTicket)
	e.GET("/filter-tickets", h.FilterTickets)
	e.POST("/tickets", h.CreateTicket)
	e.DELETE("/tickets/:id", h.DeleteTicket)
}
