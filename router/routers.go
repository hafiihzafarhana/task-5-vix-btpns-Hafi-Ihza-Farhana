package routers

import (
	"log"
	"os"
	"rakamin/projectfinal/controllers"
	"rakamin/projectfinal/database"
	"rakamin/projectfinal/middlewares"
	photopg "rakamin/projectfinal/repositories/photo_repository/photo_pg"
	userpg "rakamin/projectfinal/repositories/user_repository/user_pg"
	"rakamin/projectfinal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var PORT string

func StartApp() {
	db := database.GetPostgresInstance()

	USER_ID := os.Getenv("USER_ID")
	ADMIN_ID := os.Getenv("ADMIN_ID")

	userToInt, _ := strconv.Atoi(USER_ID)
	adminToInt, _ := strconv.Atoi(ADMIN_ID)

	userId := int(userToInt)
	adminId := int(adminToInt)

	r := gin.Default()
	route := r.Group("/api")

	userRepo := userpg.NewUserPG(db)
	userService := services.NewUserService(userRepo)
	userHandler := controllers.NewUserHandler(userService)

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)
		userRoute.POST("/login", userHandler.Login)
		userRoute.Use(middlewares.Authenticate)

		userRoute.Use(middlewares.Authorize([]int{userId, adminId}))
		userRoute.GET("/me", userHandler.GetMe)

		userRoute.Use(middlewares.Authorize(adminId))
		userRoute.PUT("/:id", userHandler.UpdateUserById)
		userRoute.DELETE("/:id", userHandler.DeleteUserById)
	}

	photoRepo := photopg.NewPhotoPG(db)
	photoService := services.NewPhotoService(photoRepo)
	photoHandler := controllers.NewPhotoHandler(photoService)

	// add photos -> post -> /photos
	// get all photos -> get -> /photos
	// update photos -> put -> /photos/:id
	// delete photos -> delete -> /photos/:id

	photoRoute := route.Group("/photos")
	{
		photoRoute.Use(middlewares.Authenticate)
		photoRoute.Use(middlewares.Authorize([]int{userId}))
		photoRoute.POST("/", photoHandler.UserCreatePhoto)
		photoRoute.PUT("/:id", photoHandler.UserUpdatePhoto)
		photoRoute.DELETE("/:id", photoHandler.UserDeletePhoto)
	}

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3000"
	}
	log.Fatalln(r.Run(":" + PORT))
}
