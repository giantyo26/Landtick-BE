package routes

import (
	"server/handlers"
	"server/pkg/middleware"
	"server/pkg/mysql"
	"server/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoute(e *echo.Group) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	ticketRepository := repositories.RepositoryTicket(mysql.DB)

	h := handlers.HandlerTransaction(transactionRepository, ticketRepository)

	e.GET("/transactions", h.FindTransactions)
	e.GET("/transactions/:id", h.GetTransaction)
	e.POST("/create-trans/:id", middleware.Auth(h.CreateTransaction))
	e.GET("/user-transactions", middleware.Auth(h.FindTransactionsByUser))
	e.GET("/user-transactions/:id", middleware.Auth(h.GetTransactionByUser))
	e.POST("/transactions/:id", middleware.Auth(h.CreateTransaction))
	e.DELETE("/transactions/:id", middleware.Auth(h.DeleteTransaction))
	e.POST("/create-trans/:id", middleware.Auth(h.CreatePayment))
	e.GET("/order-user", middleware.Auth(h.GetTransByUser))
	e.GET("/get-idpayment/:id", middleware.Auth(h.GetIdPayment))
	e.GET("/payments/:id", middleware.Auth(h.PaymentTransaction))
	e.POST("/notification", h.Notification)
}
