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

type MenuController interface {
	All(context *gin.Context)
	FindMenuByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type menuController struct {
	menuService service.MenuService
	jwtService  service.JWTService
}

func NewMenuController(menuServ service.MenuService, jwtServ service.JWTService) MenuController {
	return &menuController{
		menuService: menuServ,
		jwtService:  jwtServ,
	}
}

func (c *menuController) All(context *gin.Context) {
	var menus []model.Menu = c.menuService.All()
	res := helper.BuildResponse(true, "OK", menus)
	context.JSON(http.StatusOK, res)
}

func (c *menuController) FindMenuByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var menu model.Menu = c.menuService.FindMenuByID(id)
	if (menu == model.Menu{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		res := helper.BuildResponse(true, "OK", menu)
		context.JSON(http.StatusOK, res)
	}
}

func (c *menuController) Insert(context *gin.Context) {
	var menuCreateDTO dto.MenuCreateDTO
	errDTO := context.ShouldBind(&menuCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		customerID := c.getCustomerIDByToken(authHeader)
		convertedCustomerID, err := strconv.ParseUint(customerID, 0, 0)
		if err == nil {
			menuCreateDTO.CustomerID = convertedCustomerID
		}
		result := c.menuService.Insert(menuCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *menuController) Update(context *gin.Context) {
	var menuUpdateDTO dto.MenuUpdateDTO
	errDTO := context.ShouldBind(&menuUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} 
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	customerID := fmt.Sprintf("%v", claims["customer_id"])
	if c.menuService.IsAllowedToEdit(customerID, menuUpdateDTO.ID) {
		id, errID := strconv.ParseUint(customerID, 10, 64)
		if errID == nil {
			menuUpdateDTO.CustomerID = id
		}
		result := c.menuService.Update(menuUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		res := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusForbidden, res)
	}
}

func (c *menuController) Delete(context *gin.Context) {
	var menu model.Menu
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	} 
	menu.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	menuID := fmt.Sprintf("%v", claims["customer_id"])
	if c.menuService.IsAllowedToEdit(menuID, menu.ID) {
		c.menuService.Delete(menu)
		response := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *menuController) getCustomerIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["customer_id"])
	return id
}