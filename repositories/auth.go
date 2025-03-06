package repositories

import (
	"rantaujoeang-app-backend/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	FindUser(username string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (ar *authRepository) FindUser(username string) (models.User, error) {
	var user models.User

	err := ar.db.Where("user_username = ?", username).First(&user).Error

	return user, err
}

func (ar *authRepository) CreateUser(user models.User) (models.User, error) {
	err := ar.db.Create(&user).Error

	return user, err
}
