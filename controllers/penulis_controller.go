package controllers

import (
	"apimandiri/models"
	"apimandiri/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PenulisController interface {
	CreatePenulis(ctx *gin.Context)
	GetAllPenulis(ctx *gin.Context)
	GetPenulisByID(ctx *gin.Context)
	UpdatePenulis(ctx *gin.Context)
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

func (c *penulisControllerImpl) GetAllPenulis(ctx *gin.Context) {
	penulis, err := c.service.GetAllPenulis()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penulis)
}

func (c *penulisControllerImpl) GetPenulisByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	penulis, err := c.service.GetPenulisByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "penulis tidak tersedia"})
		return
	}
	ctx.JSON(http.StatusOK, penulis)
}

func (c *penulisControllerImpl) UpdatePenulis(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var penulis models.Penulis
	if err := ctx.ShouldBindJSON(&penulis); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	penulis.ID = uint(id)
	if err := c.service.UpdatePenulis(penulis); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"pesan": "penulis berhasil diupdate"})
}
