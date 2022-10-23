package controller

import (
	"fmt"
	"net/http"
	"rumah-makan/dto"
	"rumah-makan/helper"
	"rumah-makan/model"
	"rumah-makan/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuController interface {
	All(context *gin.Context)
	FindMenuByID(context *gin.Context)
	Insert(context *gin.Context)
	InsertMenuImage(context *gin.Context)
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


// AllMenu
// @Summary Show all menu
// @Description get all menu
// @Accept  json
// @Produce  json
// @Tags Menu
// @Success 200 {object} helper.Response
// @Router /menu [get]

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


// InsertMenu
// @Security bearerAuth
// @Summary Insert new menu
// @Description Insert new menu
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param name body string true "Name"
// @Param price body int true "Price"
// @Param description body string true "Description"
// @Param image body string true "Image"
// @Param restaurant_id body int true "Restaurant ID"
// @Success 201 {object} helper.Response
// @Router / [post]

func (c *menuController) Insert(context *gin.Context) {
	var menuCreateDTO dto.MenuCreateDTO
	errDTO := context.ShouldBind(&menuCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		createdMenu := c.menuService.Insert(menuCreateDTO)
		response := helper.BuildResponse(true, "OK", createdMenu)
		context.JSON(http.StatusCreated, response)
	}
}


// UpdateMenu
// @Security bearerAuth
// @Summary Update menu
// @Description Update menu
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param id path int true "Menu ID"
// @Param name body string true "Name"
// @Param price body int true "Price"
// @Param description body string true "Description"
// @Param image body string true "Image"
// @Param restaurant_id body int true "Restaurant ID"
// @Success 200 {object} helper.Response
// @Router /{id} [put]
 
func (c *menuController) Update(context *gin.Context) {
	var menuUpdateDTO dto.MenuUpdateDTO
	errDTO := context.ShouldBind(&menuUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	menuUpdateDTO.ID = id
	updatedMenu := c.menuService.Update(menuUpdateDTO)
	response := helper.BuildResponse(true, "OK", updatedMenu)
	context.JSON(http.StatusOK, response)
	
	
}

func (c *menuController) Delete(context *gin.Context) {
	var menu model.Menu
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	} 
	menu.ID = id
	c.menuService.Delete(menu)
	response := helper.BuildResponse(true, "OK", helper.EmptyObj{})
	context.JSON(http.StatusOK, response)
}

func (c *menuController) InsertMenuImage(context *gin.Context) {
	file, err := context.FormFile("image")
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request image file", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	menuID, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request menu id", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	path := fmt.Sprintf("images/%d-%s", menuID, file.Filename)
	err = context.SaveUploadedFile(file, path)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request save image", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var menu model.Menu = c.menuService.FindMenuByID(menuID)
	if (menu == model.Menu{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	menu.Image = path
	updatedMenu := c.menuService.InsertImage(menuID, menu.Image)
	response := helper.BuildResponse(true, "OK", updatedMenu)
	context.JSON(http.StatusOK, response)
}
