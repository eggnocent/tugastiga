package services

import (
	"apimandiri/models"
	"apimandiri/repositories"
)

type PenulisService interface {
	CreatePenulis(penulis models.Penulis) error
	GetAllPenulis() ([]models.Penulis, error)
	GetPenulisByID(id uint) (models.Penulis, error)
	UpdatePenulis(penulis models.Penulis) error
}

type penulisServiceImpl struct {
	repo repositories.PenulisRepository
}

func NewPenulisService(repo repositories.PenulisRepository) PenulisService {
	return &penulisServiceImpl{repo}
}

func (s *penulisServiceImpl) CreatePenulis(penulis models.Penulis) error {
	return s.repo.CreatePenulis(penulis)
}

func (s *penulisServiceImpl) GetAllPenulis() ([]models.Penulis, error) {
	return s.repo.GetAllPenulis()
}

func (s *penulisServiceImpl) GetPenulisByID(id uint) (models.Penulis, error) {
	return s.repo.GetPenulisByID(id)
}

func (s *penulisServiceImpl) UpdatePenulis(penulis models.Penulis) error {
	return s.repo.UpdatePenulis(penulis)
}
