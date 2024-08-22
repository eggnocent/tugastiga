package controllers

import (
	"apimandiri/models"
	"apimandiri/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BukuController interface {
	AddUserBook(ctx *gin.Context)
	GetUserBook(ctx *gin.Context)
	UpdateUserBook(ctx *gin.Context) // Ditambahkan
	DeleteUserBook(ctx *gin.Context) // Ditambahkan
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
		"Penulis":   buku.Penulis,
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
