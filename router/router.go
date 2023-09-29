package router

import (
	"task-5-pbi-btpns-arufhakim/controllers"

	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	router := gin.Default()

	authController := controllers.AuthController{}
	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)

	userController := controllers.UserController{}
	router.PUT("/users/:userId", userController.Update)
	router.DELETE("/users/:userId", userController.Delete)

	photoController := controllers.PhotoController{}
	router.GET("/photos", photoController.Get)
	router.POST("/photos", photoController.Create)
	router.PUT("/photos/:photoId", photoController.Update)
	router.DELETE("/photos/:photoId", photoController.Delete)

	return router
}
