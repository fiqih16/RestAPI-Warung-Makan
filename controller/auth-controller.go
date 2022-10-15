package controller

import "github.com/gin-gonic/gin"

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
