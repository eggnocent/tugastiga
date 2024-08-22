package repositories

import (
	"apimandiri/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBooks() ([]models.Buku, error)
	GetBookByUserID(userID uint) (*models.Buku, error)
	GetBookByID(id uint) (*models.Buku, error)
	AddBookToUser(book models.Buku) error
	UpdateBook(book models.Buku) error
	DeleteBook(userID uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) GetAllBooks() ([]models.Buku, error) {
	var books []models.Buku
	err := r.db.Preload("Penulis").Preload("User").Find(&books).Error
	return books, err
}

func (r *bookRepository) GetBookByID(id uint) (*models.Buku, error) {
	var book models.Buku
	err := r.db.Preload("Penulis").Preload("User").First(&book, id).Error
	return &book, err
}

func (r *bookRepository) GetBookByUserID(userID uint) (*models.Buku, error) {
	var book models.Buku
	err := r.db.Preload("User").Preload("Penulis").Where("user_id = ?", userID).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) AddBookToUser(book models.Buku) error {
	return r.db.Create(&book).Error
}

func (r *bookRepository) UpdateBook(book models.Buku) error {
	return r.db.Save(&book).Error
}

func (r *bookRepository) DeleteBook(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Buku{}).Error
}
