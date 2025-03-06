package repositories

import (
	"rantaujoeang-app-backend/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	FindProfile(username string) (models.User, error)
	UpdateProfile(user models.User) (models.User, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *profileRepository {
	return &profileRepository{db}
}

func (pr *profileRepository) FindProfile(username string) (models.User, error) {
	var profile models.User

	err := pr.db.Where("user_username = ?", username).First(&profile).Error

	return profile, err
}

func (pr *profileRepository) UpdateProfile(user models.User) (models.User, error) {
	err := pr.db.Save(&user).Error

	return user, err
}
