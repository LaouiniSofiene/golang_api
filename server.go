package main

import (
	"github.com/LaouiniSofiene/golang_api/config"
	"github.com/LaouiniSofiene/golang_api/controller"
	"github.com/LaouiniSofiene/golang_api/repository"
	"github.com/LaouiniSofiene/golang_api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	userRoutes := r.Group("api/user")
	{
		userRoutes.GET("/", userController.All)
		userRoutes.POST("/", userController.Insert)
		userRoutes.GET("/:id", userController.FindByID)
		userRoutes.PUT("/:id", userController.Update)
		userRoutes.DELETE("/:id", userController.Delete)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
