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

	menuRoutes := r.Group("api/menu", middleware.AuthorizeJWT(jwtService))
	{
		menuRoutes.GET("/", menuController.All)
		menuRoutes.POST("/", menuController.Insert)
		menuRoutes.GET("/:id", menuController.FindMenuByID)
		menuRoutes.PUT("/:id", menuController.Update)
		menuRoutes.DELETE("/:id", menuController.Delete)
	}

	transactionRoutes := r.Group("api/transaction", middleware.AuthorizeJWT(jwtService))
	{
		transactionRoutes.POST("/", transactionController.InsertTransaction)
	}
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}