package service

import (
	"log"
	"rumah-makan/dto"
	"rumah-makan/model"
	"rumah-makan/repository"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateCustomer(customer dto.RegisterDto) model.Customer
	FindByEmail(email string) model.Customer
	IsDuplicateEmail(email string) bool
}

type authService struct {
	customerRepository repository.CustomerRepository
}

func NewAuthService(customerRepo repository.CustomerRepository) AuthService {
	return &authService{
		customerRepository: customerRepo,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.customerRepository.VerifyCredential(email, password)
	if v, ok := res.(model.Customer); ok {
		comparedPassword := comparedPassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return nil
}

func (service *authService) CreateCustomer(customer dto.RegisterDto) model.Customer {
	customerToCreate := model.Customer{}
	err := smapping.FillStruct(&customerToCreate, smapping.MapFields(&customer))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.customerRepository.InsertCustomer(customerToCreate)
	return res
}

func (service *authService) FindByEmail(email string) model.Customer {
	return service.customerRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.customerRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparedPassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}