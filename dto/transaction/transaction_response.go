package transactiondto

import "server/models"

type TransactionResponse struct {
	ID            int           `json:"id"`
	TransactionID int           `json:"transaction_id" `
	UserID        int           `json:"user_id"`
	User          models.User   `json:"user"`
	TicketID      int           `json:"ticket_id"`
	Ticket        models.Ticket `json:"ticket"`
	Adult         int           `json:"adult"`
	Infant        int           `json:"infant"`
	Qty           int           `json:"qty"`
	Total         int           `json:"total"`
	Status        string        `json:"status" gorm:"type: VARCHAR(25)"`
}
