package main

import (
	"apimandiri/config"
	"apimandiri/controllers"
	"apimandiri/models"
	"apimandiri/repositories"
	"apimandiri/server"
	"apimandiri/services"
)

func main() {
	db := config.InitDB()
	db.AutoMigrate(&models.User{}, &models.Buku{})

	// Inisialisasi repository, service, dan controller
	userRepo := repositories.NewUserRepository(db)               // Mengambil db sebagai parameter untuk mengakses database.
	userService := services.NewUserService(userRepo)             // UserService: Menggunakan UserRepository untuk operasi logika bisnis.
	userController := controllers.NewUserController(userService) // UserController: Menggunakan UserService untuk menangani request HTTP.
	authService := services.NewAuthService(userRepo)             // AuthService: Menggunakan UserRepository untuk autentikasi.
	authController := controllers.NewAuthController(authService) // AuthController: Menggunakan AuthService untuk login/logout.

	// Inisialisasi router dan jalankan server
	r := server.InitRouter(authController, userController)
	r.Run() // Menjalankan server pada port default (8080)
}
