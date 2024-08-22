package server

import (
	"apimandiri/controllers"
	"apimandiri/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter(authController controllers.AuthController, userController controllers.UserController, bukuController controllers.BukuController, penulisController controllers.PenulisController) *gin.Engine {
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

	// Rute untuk buku
	protected.POST("/users/:id/buku", bukuController.AddUserBook)
	protected.GET("/users/:id/buku", bukuController.GetUserBook)
	protected.PUT("/users/:id/buku", bukuController.UpdateUserBook)
	protected.DELETE("/users/:id/buku", bukuController.DeleteUserBook)

	protected.POST("/penulis", penulisController.CreatePenulis)
	protected.GET("/penulis", penulisController.GetAllPenulis)
	return r
}
