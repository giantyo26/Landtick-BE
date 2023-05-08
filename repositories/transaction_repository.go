package repositories

import (
	"server/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	FindTransactionsByUser(UserID int) ([]models.Transaction, error)
	GetTransactionByUser(UserID, TransactionID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTicketById(ID int) (models.Ticket, error)
	CreatePayment(transaction models.Transaction) (models.Transaction, error)
	Payment(payment models.Transaction) (models.Transaction, error)
	GetPaymentByIdTrans(ID int) (models.Transaction, error)
	GetTicketTransaction(UserID int) ([]models.Transaction, error)
	UpdatePayment(status string, ID int) (models.Transaction, error)
	UpdateTransaction(status string, ID int) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("User").Preload("Ticket").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Ticket").First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) FindTransactionsByUser(UserID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("User").Preload("Ticket").Where("user_id = ?", UserID).Error

	return transaction, err
}

func (r *repository) GetTransactionByUser(UserID, TransactionID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Ticket").Where("user_id = ? AND id = ?", UserID, TransactionID).First(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransUser(UserID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Ticket.StartStation").Preload("Ticket.DestinationStation").Where("user = ?", UserID).Where("status = ?").Find(&transaction).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Preload("Ticket").Create(&transaction).Error

	return transaction, err
}

func (r *repository) GetTicketById(ID int) (models.Ticket, error) {
	var ticket models.Ticket
	err := r.db.Preload("StartStation").Preload("DestinationStation").First(&ticket, "id = ?", ID).Error

	return ticket, err
}

func (r *repository) GetTicketTransaction(UserID int) ([]models.Transaction, error) {
	var transaction []models.Transaction

	err := r.db.Preload("User").Preload("Ticket.StartStation").Preload("Ticket.DestinationStation").Where("user_id = ?", UserID).Find(&transaction).Error

	return transaction, err
}

func (r *repository) CreatePayment(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Ticket.StartStation").Preload("Ticket.DestinationStation").Create(&transaction).Error

	return transaction, err
}

func (r *repository) GetPaymentByIdTrans(ID int) (models.Transaction, error) {
	var payment models.Transaction
	err := r.db.Preload("Ticket").Preload("User").Where("id = ?", ID).Find(&payment).Error

	return payment, err
}

func (r *repository) Payment(payment models.Transaction) (models.Transaction, error) {
	err := r.db.Save(&payment).Error
	return payment, err
}

func (r *repository) UpdatePayment(status string, ID int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("Ticket").First(&transaction, ID)

	if status != transaction.Status && status == "success" {
		var ticket models.Ticket
		r.db.First(&ticket, transaction.Ticket.ID)
		ticket.Qty = ticket.Qty - 1
		r.db.Save(&ticket)
	}

	transaction.Status = status
	error := r.db.Save(&transaction).Error
	return transaction, error
}

func (r *repository) UpdateTransaction(status string, ID int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("User").Preload("Ticket").First(&transaction, ID)

	if status != transaction.Status && status == "success" {
		var ticket models.Ticket
		r.db.First(&ticket, transaction.Ticket.ID)
		ticket.Qty = ticket.Qty - transaction.Adult - transaction.Infant
		r.db.Save(&ticket)
	}

	transaction.Status = status
	error := r.db.Save(&transaction).Error
	return transaction, error
}

func (r *repository) DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&transaction, ID).Scan(&transaction).Error

	return transaction, err
}
