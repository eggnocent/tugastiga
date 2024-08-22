package repositories

import (
	"apimandiri/models"

	"gorm.io/gorm"
)

type PenulisRepository interface {
	CreatePenulis(penulis models.Penulis) error
	GetAllPenulis() ([]models.Penulis, error)
}

type penulisRepositoryImpl struct {
	db *gorm.DB
}

func NewPenulisRepository(db *gorm.DB) PenulisRepository {
	return &penulisRepositoryImpl{db}
}

func (r *penulisRepositoryImpl) CreatePenulis(penulis models.Penulis) error {
	return r.db.Create(&penulis).Error
}

func (r *penulisRepositoryImpl) GetAllPenulis() ([]models.Penulis, error) {
	var penulis []models.Penulis
	err := r.db.Find(&penulis).Error
	return penulis, err
}
