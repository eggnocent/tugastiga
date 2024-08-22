package controllers

import (
	"apimandiri/models"
	"apimandiri/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BukuController interface {
	AddUserBook(ctx *gin.Context)
	GetAllBooks(ctx *gin.Context)
	GetUserBook(ctx *gin.Context)
	GetBookByID(ctx *gin.Context)
	UpdateUserBook(ctx *gin.Context)
	DeleteUserBook(ctx *gin.Context)
}

type bukuController struct {
	service services.BookService
}

func NewBukuController(service services.BookService) BukuController {
	return &bukuController{service}
}

func (c *bukuController) AddUserBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var buku models.Buku
	if err := ctx.ShouldBindJSON(&buku); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.service.AddBookToUser(id, buku)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Book added to user"})
}

func (c *bukuController) GetAllBooks(ctx *gin.Context) {
	// Panggil service untuk mendapatkan semua buku
	books, err := c.service.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	// Buat response JSON manual untuk setiap buku
	var booksResponse []gin.H
	for _, book := range books {
		bookResponse := gin.H{
			"ID":        book.ID,
			"NamaBuku":  book.NamaBuku,
			"Penulis":   book.Penulis.NamaPenulis,
			"TglTerbit": book.TglTerbit,
			"User":      book.User.Username,
			"UserID":    book.UserID,
		}
		booksResponse = append(booksResponse, bookResponse)
	}

	// Kirim response JSON
	ctx.JSON(http.StatusOK, booksResponse)
}

func (c *bukuController) GetBookByID(ctx *gin.Context) {
	// Ambil ID dari URL parameter dan konversi ke uint
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	// Panggil service untuk mendapatkan data buku berdasarkan ID
	book, err := c.service.GetBookByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Buat response manual seperti yang Anda inginkan
	response := gin.H{
		"ID":        book.ID,
		"NamaBuku":  book.NamaBuku,
		"Penulis":   book.Penulis.NamaPenulis,
		"TglTerbit": book.TglTerbit,
		"User":      book.User.Username,
		"UserID":    book.UserID,
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *bukuController) GetUserBook(ctx *gin.Context) {
	id := ctx.Param("id")
	buku, err := c.service.GetBookByUserID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ada"})
		return
	}

	response := map[string]interface{}{
		"ID":        buku.ID,
		"NamaBuku":  buku.NamaBuku,
		"Penulis":   buku.Penulis.NamaPenulis,
		"TglTerbit": buku.TglTerbit,
		"UserID":    buku.UserID,
		"User":      buku.User.Username,
	}

	ctx.JSON(http.StatusOK, response)
}

// Fungsi untuk mengupdate buku milik user
func (c *bukuController) UpdateUserBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var buku models.Buku
	if err := ctx.ShouldBindJSON(&buku); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.service.UpdateBook(id, buku)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

// Fungsi untuk menghapus buku milik user
func (c *bukuController) DeleteUserBook(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.service.DeleteBook(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
