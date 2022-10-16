package main

import (
	"rumah-makan/config"
	"rumah-makan/controller"
	"rumah-makan/middleware"
	"rumah-makan/repository"
	"rumah-makan/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetDBConn()
	// Repository
	customerRepository repository.CustomerRepository = repository.NewCustomerRepository(db)
	
	// Service
	jwtService service.JWTService = service.NewJWTService()
	authService service.AuthService = service.NewAuthService(customerRepository)
	customerService service.CustomerService = service.NewCustomerService(customerRepository)

	// Controller
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	customerController controller.CustomerController = controller.NewCustomerController(customerService, jwtService)
)

func main() {
	defer config.CloseDBConn(db)
	r := gin.Default()
	
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	customerRoutes := r.Group("api/customer", middleware.AuthorizeJWT(jwtService))
	{
		customerRoutes.GET("/profile", customerController.Profile)
		customerRoutes.PUT("/profile", customerController.Update)
	}
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}