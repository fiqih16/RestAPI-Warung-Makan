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

type CustomerController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type customerController struct {
	customerService service.CustomerService
	jwtService      service.JWTService
}

func NewCustomerController(customerService service.CustomerService, jwtService service.JWTService) CustomerController {
	return &customerController{
		customerService: customerService,
		jwtService:      jwtService,
	}
}

func (c *customerController) Update(context *gin.Context) {
	var customerUpdateDTO dto.CustomerUpdateDTO
	errDTO := context.ShouldBind(&customerUpdateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["customer_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	customerUpdateDTO.ID = id
	u := c.customerService.Update(customerUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (c *customerController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["customer_id"])
	customer := c.customerService.Profile(id)
	res := helper.BuildResponse(true, "OK!", customer)
	context.JSON(http.StatusOK, res)
}