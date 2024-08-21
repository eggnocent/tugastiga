package controllers

import (
	"apimandiri/models"
	"apimandiri/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}

type userController struct {
	service services.UserService
}

func NewUserController(services services.UserService) UserController {
	return &userController{services}
}

func (c *userController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}
	if user.FullName == "" || user.Username == "" || user.Password == "" || user.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "harus mengisi semua kolom"})
		return
	}
	if err := c.service.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errpr": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"pesan": "User Berhasil DIbuat"})
}

func (c *userController) GetAllUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *userController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.service.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	_, err = c.service.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "ID tidak ditemukan"})
		return
	}
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = uint(userID)

	if err := c.service.UpdateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"pesan": "User telah di update"})
}
