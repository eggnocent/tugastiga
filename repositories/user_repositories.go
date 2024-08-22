package repositories

import (
	"apimandiri/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) error
	FindAll(users *[]models.User) error
	FindByID(id uint, user *models.User) error
	Update(user models.User) error
	FindByUsername(username string, user *models.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user models.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) FindAll(users *[]models.User) error {
	return r.db.Find(users).Error
}

func (r *userRepository) FindByID(id uint, user *models.User) error {
	return r.db.First(user, id).Error
}

func (r *userRepository) Update(user models.User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) FindByUsername(username string, user *models.User) error {
	return r.db.Where("username = ?", username).First(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Unscoped().Delete(&models.User{}, id).Error
}
