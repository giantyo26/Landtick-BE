package models

type Transaction struct {
	ID            int    `json:"id" gorm:"primaryKey;autoIncrement"`
	TransactionID int    `json:"transaction_id"`
	UserID        int    `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User          User   `json:"user" `
	TicketID      int    `json:"ticket_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Ticket        Ticket `json:"ticket" `
	Status        string `json:"status"`
	Adult         int    `json:"adult" `
	Infant        int    `json:"infant"`
	Qty           int    `json:"qty"`
	Total         int    `json:"total" form:"total"`
}
