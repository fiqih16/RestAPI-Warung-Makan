package controller

import (
	"fmt"
	"net/http"
	"rumah-makan/dto"
	"rumah-makan/helper"
	"rumah-makan/model"
	"rumah-makan/service"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	InsertTransaction(context *gin.Context)
	UpdateTransaction(context *gin.Context)
	DeleteTransaction(context *gin.Context)
	AllTransaction(context *gin.Context)
}

type transactionController struct {
	transactionService service.TransactionService
	jwtService         service.JWTService
}

func NewTransactionController(transactionServ service.TransactionService, jwtServ service.JWTService) TransactionController {
	return &transactionController{
		transactionService: transactionServ,
		jwtService:         jwtServ,
	}
}

func (c *transactionController) InsertTransaction(context *gin.Context) {
	var transactionCreateDTO dto.TransactionCreateDTO
	errDTO := context.ShouldBind(&transactionCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		token := context.GetHeader("Authorization")
		customerID := c.getCustomerIDByToken(token)
		convertedCustomerID, err := strconv.ParseUint(customerID, 0, 0)
		if err == nil {
			transactionCreateDTO.CustomerID = convertedCustomerID
		}
		result := c.transactionService.Insert(transactionCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *transactionController) UpdateTransaction(context *gin.Context) {
	var transactionUpdateDTO dto.TransactionUpdateDTO
	errDTO := context.ShouldBind(&transactionUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		token := context.GetHeader("Authorization")
		customerID := c.getCustomerIDByToken(token)
		convertedCustomerID, err := strconv.ParseUint(customerID, 0, 0)
		if err == nil {
			transactionUpdateDTO.CustomerID = convertedCustomerID
		}
		result := c.transactionService.Update(transactionUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} 
}

func (c *transactionController) DeleteTransaction(context *gin.Context) {
	var transaction model.Transaction
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		transaction.ID = id
		c.transactionService.Delete(transaction)
		res := helper.BuildResponse(true, "OK", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	}
}

func (c *transactionController) AllTransaction(context *gin.Context) {
	var transactions []model.Transaction = c.transactionService.All()
	res := helper.BuildResponse(true, "OK", transactions)
	context.JSON(http.StatusOK, res)
}

func (c *transactionController) getCustomerIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["customer_id"])
	return id
}