package controllers

import (
	"apimandiri/models"
	"apimandiri/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PenulisController interface {
	CreatePenulis(ctx *gin.Context)
}

type penulisControllerImpl struct {
	service services.PenulisService
}

func NewPenulisController(service services.PenulisService) PenulisController {
	return &penulisControllerImpl{service}
}

func (c *penulisControllerImpl) CreatePenulis(ctx *gin.Context) {
	var penulis models.Penulis
	if err := ctx.ShouldBindJSON(&penulis); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreatePenulis(penulis); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"pesan": "Penulis berhasil dibuat"})
}
