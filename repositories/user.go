package repositories

import (
	"rantaujoeang-app-backend/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	FindUser(username string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindUsers() ([]models.User, error) {
	var users []models.User

	err := ur.db.Find(&users).Error

	return users, err
}

func (ur *userRepository) FindUser(username string) (models.User, error) {
	var user models.User

	err := ur.db.Where("user_username = ?", username).First(&user).Error

	return user, err
}

func (ur *userRepository) CreateUser(user models.User) (models.User, error) {
	err := ur.db.Create(&user).Error

	return user, err
}

func (ur *userRepository) UpdateUser(user models.User) (models.User, error) {
	err := ur.db.Save(&user).Error

	return user, err
}

func (ur *userRepository) DeleteUser(user models.User) (models.User, error) {
	err := ur.db.Delete(&user).Error

	return user, err
}
