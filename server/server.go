package server

import (
	"apimandiri/controllers"
	"apimandiri/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter(authController controllers.AuthController, userController controllers.UserController) *gin.Engine {
	r := gin.Default()

	r.POST("/login", authController.Login)
	r.POST("/logout", authController.Logout)

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	protected.POST("/users", userController.CreateUser)
	protected.GET("/users", userController.GetAllUsers)
	protected.GET("/users/:id", userController.GetUserByID)
	protected.PUT("/users/:id", userController.UpdateUser)
	protected.DELETE("/users/:id", userController.DeleteUser)
	protected.POST("/users/:id/buku", userController.AddUserBook)
	protected.GET("/users/:id/buku", userController.GetUserBook)
	return r
}
