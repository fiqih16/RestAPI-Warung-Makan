package controller

import (
	"fmt"
	"net/http"
	"rumah-makan/dto"
	"rumah-makan/helper"
	"rumah-makan/service"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	InsertTransaction(context *gin.Context)
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

func (c *transactionController) getCustomerIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["customer_id"])
	return id
}