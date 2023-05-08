package routes

import (
	"server/handlers"
	"server/pkg/mysql"
	"server/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	UserRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(UserRepository)

	e.GET("/users", h.FindUsers)
	e.GET("/users/:id", h.GetUser)
	e.DELETE("/users/:id", h.DeleteUser)
}
