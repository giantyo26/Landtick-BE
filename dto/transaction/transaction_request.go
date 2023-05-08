package transactiondto

type TransactionRequest struct {
	UserID   int    `json:"user_id" `
	TicketId int    `json:"ticket_id"`
	Adult    int    `json:"adult"`
	Infant   int    `json:"infant"`
	Qty      int    `json:"qty"`
	Total    int    `json:"total"`
	Status   string `json:"status"`
}
