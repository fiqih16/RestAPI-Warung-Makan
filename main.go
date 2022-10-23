package main

import (
	"rumah-makan/config"
	"rumah-makan/controller"
	"rumah-makan/docs"
	_ "rumah-makan/docs"
	"rumah-makan/repository"
	"rumah-makan/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetDBConn()
	// Repository
	customerRepository repository.CustomerRepository = repository.NewCustomerRepository(db)
	menuRepository repository.MenuRepository = repository.NewMenuRepository(db)
	transactionRepository repository.TransactionRepository = repository.NewTransactionRepository(db)

	// Service
	jwtService service.JWTService = service.NewJWTService()
	authService service.AuthService = service.NewAuthService(customerRepository)
	customerService service.CustomerService = service.NewCustomerService(customerRepository)
	menuService service.MenuService = service.NewMenuService(menuRepository)
	transactionService service.TransactionService = service.NewTransactionService(transactionRepository, menuRepository)

	// Controller
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	customerController controller.CustomerController = controller.NewCustomerController(customerService, jwtService)
	menuController controller.MenuController = controller.NewMenuController(menuService, jwtService)
	transactionController controller.TransactionController = controller.NewTransactionController(transactionService, jwtService)
)

// @title Rumah Makan API Documentation
// @version 1.0
// @description This is a sample server for a Rumah Makan API.
// @termsOfService http://swagger.io/terms/

// @contact.name Swagger API Team
// @contact.url http://swagger.io
// @contact.email

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/menu

func main() {
	defer config.CloseDBConn(db)
	r := gin.Default()
	
	// authRoutes := r.Group("api/auth")
	// {
	// 	authRoutes.POST("/login", authController.Login)
	// 	authRoutes.POST("/register", authController.Register)
	// }

	// customerRoutes := r.Group("api/customer", middleware.AuthorizeJWT(jwtService))
	// {
	// 	customerRoutes.GET("/profile", customerController.Profile)
	// 	customerRoutes.PUT("/profile", customerController.Update)
	// }

	menuRoutes := r.Group("api/menu")
	{
		menuRoutes.GET("/", menuController.All)
		menuRoutes.POST("/", menuController.Insert)
		menuRoutes.GET("/:id", menuController.FindMenuByID)
		menuRoutes.PUT("/:id", menuController.Update)
		menuRoutes.DELETE("/:id", menuController.Delete)
		menuRoutes.POST("/image/:id", menuController.InsertMenuImage)
	}

	// transactionRoutes := r.Group("api/transaction", middleware.AuthorizeJWT(jwtService))
	// {
	// 	transactionRoutes.POST("/", transactionController.InsertTransaction)
	// 	transactionRoutes.GET("/", transactionController.AllTransaction)
	// 	transactionRoutes.PUT("/:id", transactionController.UpdateTransaction)
	// 	transactionRoutes.DELETE("/:id", transactionController.DeleteTransaction)
	// }

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/menu/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}