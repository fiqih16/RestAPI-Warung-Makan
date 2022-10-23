package controller

import (
	"net/http"
	"rumah-makan/dto"
	"rumah-makan/helper"
	"rumah-makan/model"
	"rumah-makan/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	// masukkan service yang kalian butuh
	authService service.AuthService
	jwtService service.JWTService
}

// NewAuthController membuat instance baru dari AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService: jwtService,
	}
}


// Login
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} model.Customer
// @Failure 400 {object} helper.APIResponse
// @Router /auth/login [post] 

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDto
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(model.Customer); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}


// Register is a function to register new user
// @Summary Register new user
// @Description Register new user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param register body dto.RegisterDto true "Register"
// @Success 201 {object} dto.RegisterDto "Register"
// @Failure 400 {object} helper.APIResponse "Error"
// @Router /register [post]
// @Security ApiKeyAuth

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDto
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		
	} else {
		createdCustomer := c.authService.CreateCustomer(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdCustomer.ID, 10))
		createdCustomer.Token = token
		response := helper.BuildResponse(true, "OK!", createdCustomer)
		ctx.JSON(http.StatusCreated, response)
	} 
}