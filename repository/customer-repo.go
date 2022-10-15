package repository

import (
	"log"
	"rumah-makan/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	InsertCustomer(customer model.Customer) model.Customer
	UpdateCustomer(customer model.Customer) model.Customer
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) model.Customer
}

type customerConnection struct {
	connection *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerConnection{
		connection: db,
	}
}

func (db *customerConnection) InsertCustomer(customer model.Customer) model.Customer {
	customer.Password = hashAndSalt([]byte(customer.Password))
	db.connection.Save(&customer)
	return customer
}

func (db *customerConnection) UpdateCustomer(customer model.Customer) model.Customer {
	if customer.Password != "" {
		customer.Password = hashAndSalt([]byte(customer.Password))
	} else {
		var tempCustomer model.Customer
		db.connection.Find(&tempCustomer, customer.ID)
		customer.Password = tempCustomer.Password
	}
	db.connection.Save(&customer)
	return customer
}

func (db *customerConnection) VerifyCredential(email string, password string) interface{} {
	var customer model.Customer
	res := db.connection.Where("email = ?", email).Take(&customer)
	if res.Error == nil {
		return customer
	}
	return nil
}

func (db *customerConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var customer model.Customer
	return db.connection.Where("email = ?", email).Take(&customer)
}

func (db *customerConnection) FindByEmail(email string) model.Customer {
	var customer model.Customer
	db.connection.Where("email = ?", email).Take(&customer)
	return customer
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}