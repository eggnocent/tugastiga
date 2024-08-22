package services

import (
	"apimandiri/models"
	"apimandiri/repositories"
	"errors"
	"strconv"
)

type BookService interface {
	GetBookByUserID(userID string) (*models.Buku, error)
	AddBookToUser(userID string, book models.Buku) error
	UpdateBook(userID string, book models.Buku) error
	DeleteBook(userID string) error
}

type bookServiceImpl struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookServiceImpl{repo}
}

func (s *bookServiceImpl) GetBookByUserID(userID string) (*models.Buku, error) {
	id, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}
	return s.repo.GetBookByUserID(uint(id))
}

func (s *bookServiceImpl) AddBookToUser(userID string, book models.Buku) error {
	id, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return errors.New("invalid user ID")
	}
	book.UserID = uint(id)
	return s.repo.AddBookToUser(book)
}

func (s *bookServiceImpl) UpdateBook(userID string, book models.Buku) error {
	id, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return errors.New("invalid user ID")
	}
	existingBook, err := s.repo.GetBookByUserID(uint(id))
	if err != nil {
		return err
	}
	book.ID = existingBook.ID
	book.UserID = existingBook.UserID
	return s.repo.UpdateBook(book)
}

func (s *bookServiceImpl) DeleteBook(userID string) error {
	id, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return errors.New("invalid user ID")
	}
	return s.repo.DeleteBook(uint(id))
}
