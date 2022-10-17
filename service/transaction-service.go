package service

import (
	"rumah-makan/dto"
	"rumah-makan/model"
	"rumah-makan/repository"
	"time"
)


type TransactionService interface {
	Insert(t dto.TransactionCreateDTO) model.Transaction
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
	menuRepository 	  repository.MenuRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository, menuRepo repository.MenuRepository) TransactionService {
	return &transactionService{
		transactionRepository: transactionRepo,
		menuRepository: menuRepo,
	}
}

func (service *transactionService) Insert(t dto.TransactionCreateDTO) model.Transaction {
	transaction := model.Transaction{}
	transaction.CustomerID = t.CustomerID
	transaction.MenuID = t.MenuID
	transaction.JumlahBeli = t.JumlahBeli
	transaction.TotalBayar = service.menuRepository.FindMenuByID(t.MenuID).Harga * t.JumlahBeli 
	transaction.Tanggal = time.Now()

	res := service.transactionRepository.InsertTransaction(transaction)
	return res
}
