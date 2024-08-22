package main

import (
	"apimandiri/config"
	"apimandiri/controllers"
	"apimandiri/models"
	"apimandiri/repositories"
	"apimandiri/server"
	"apimandiri/services"
	"log"
)

func main() {
	db := config.InitDB()
	if err := db.AutoMigrate(&models.User{}, &models.Penulis{}, &models.Buku{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// Inisialisasi repository, service, dan controller
	userRepo := repositories.NewUserRepository(db)               // Mengambil db sebagai parameter untuk mengakses database.
	userService := services.NewUserService(userRepo)             // UserService: Menggunakan UserRepository untuk operasi logika bisnis.
	userController := controllers.NewUserController(userService) // UserController: Menggunakan UserService untuk menangani request HTTP.
	authService := services.NewAuthService(userRepo)             // AuthService: Menggunakan UserRepository untuk autentikasi.
	authController := controllers.NewAuthController(authService) // AuthController: Menggunakan AuthService untuk login/logout.

	bookRepo := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	bukuController := controllers.NewBukuController(bookService)

	penulisRepo := repositories.NewPenulisRepository(db)
	penulisService := services.NewPenulisService(penulisRepo)
	penulisController := controllers.NewPenulisController(penulisService)

	// Inisialisasi router dan jalankan server
	r := server.InitRouter(authController, userController, bukuController, penulisController) // Tambahkan bukuController sebagai parameter
	r.Run()                                                                                   // Menjalankan server pada port default (8080)
}
