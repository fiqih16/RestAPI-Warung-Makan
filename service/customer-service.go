package service

import (
	"log"
	"rumah-makan/dto"
	"rumah-makan/model"
	"rumah-makan/repository"

	"github.com/mashingan/smapping"
)

type CustomerService interface {
	Update(customer dto.CustomerUpdateDTO) model.Customer
	Profile(customerID string) model.Customer
}

type customerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRepository: customerRepo,
	}
}

func (service *customerService) Update(customer dto.CustomerUpdateDTO) model.Customer {
	customerToUpdate := model.Customer{}
	err := smapping.FillStruct(&customerToUpdate, smapping.MapFields(&customer))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.customerRepository.UpdateCustomer(customerToUpdate)
	return res
}

func (service *customerService) Profile(customerID string) model.Customer {
	return service.customerRepository.ProfileCustomer(customerID)
}