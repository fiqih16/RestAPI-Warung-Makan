package service

import (
	"fmt"
	"rumah-makan/dto"
	"rumah-makan/model"
	"rumah-makan/repository"
	"time"
)


type TransactionService interface {
	Insert(t dto.TransactionCreateDTO) model.Transaction
	Update(t dto.TransactionUpdateDTO) model.Transaction
	Delete(t model.Transaction)
	All() []model.Transaction
	IsAllowedToEdit(customerID string, transactionID uint64) bool
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

func (service *transactionService) Update(t dto.TransactionUpdateDTO) model.Transaction {
	transaction := model.Transaction{}
	transaction.ID = t.ID
	transaction.CustomerID = t.CustomerID
	transaction.MenuID = t.MenuID
	transaction.JumlahBeli = t.JumlahBeli
	transaction.TotalBayar = service.menuRepository.FindMenuByID(t.MenuID).Harga * t.JumlahBeli
	transaction.Tanggal = time.Now().AddDate(1, 0, 0)

	res := service.transactionRepository.UpdateTransaction(transaction)
	return res
}

func (service *transactionService) Delete(t model.Transaction) {
	service.transactionRepository.DeleteTransaction(t)
}

func (service *transactionService) All() []model.Transaction {
	return service.transactionRepository.AllTransaction()
}

func (service *transactionService) IsAllowedToEdit(customerID string, transactionID uint64) bool {
	t := service.transactionRepository.FindTransactionByID(transactionID)
	id := fmt.Sprintf("%v", t.CustomerID)
	return customerID == id
}