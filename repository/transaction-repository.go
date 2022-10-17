package repository

import (
	"rumah-makan/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	InsertTransaction(t model.Transaction) model.Transaction
	UpdateTransaction(t model.Transaction) model.Transaction
	DeleteTransaction(t model.Transaction)
	AllTransaction() []model.Transaction
	FindTransactionByID(transactionID uint64) model.Transaction
}

type transactionConnection struct {
	connection *gorm.DB
}

func NewTransactionRepository(dbConn *gorm.DB) TransactionRepository {
	return &transactionConnection{
		connection: dbConn,
	}
}

func (db *transactionConnection) InsertTransaction(t model.Transaction) model.Transaction {
	db.connection.Save(&t)
	return t
}

func (db *transactionConnection) UpdateTransaction(t model.Transaction) model.Transaction {
	db.connection.Save(&t)
	db.connection.Preload("Customer").Find(&t)
	return t
}

func (db *transactionConnection) DeleteTransaction(t model.Transaction) {
	db.connection.Delete(&t)
}

func (db *transactionConnection) AllTransaction() []model.Transaction {
	var transactions []model.Transaction
	db.connection.Preload("Customer").Find(&transactions)
	return transactions
}

func (db *transactionConnection) FindTransactionByID(transactionID uint64) model.Transaction {
	var transaction model.Transaction
	db.connection.Preload("Customer").Find(&transaction, transactionID)
	return transaction
}