package repositories

import (
	"rantaujoeang-app-backend/models"

	"gorm.io/gorm"
)

type DormRepository interface {
	FindDorms() ([]models.Dorm, error)
	FindDorm(dorm_code string) (models.Dorm, error)
	CreateDorm(dorm models.Dorm) (models.Dorm, error)
	UpdateDorm(dorm models.Dorm) (models.Dorm, error)
	DeleteDorm(dorm models.Dorm) (models.Dorm, error)
}

type dormRepository struct {
	db *gorm.DB
}

func NewDormRepository(db *gorm.DB) *dormRepository {
	return &dormRepository{db}
}

func (dr *dormRepository) FindDorms() ([]models.Dorm, error) {
	var dorms []models.Dorm

	err := dr.db.Find(&dorms).Error

	return dorms, err
}

func (dr *dormRepository) FindDorm(dorm_code string) (models.Dorm, error) {
	var dorm models.Dorm

	err := dr.db.Where("dorm_code = ?", dorm_code).First(&dorm).Error

	return dorm, err
}

func (dr *dormRepository) CreateDorm(dorm models.Dorm) (models.Dorm, error) {
	err := dr.db.Create(&dorm).Error

	return dorm, err
}

func (dr *dormRepository) UpdateDorm(dorm models.Dorm) (models.Dorm, error) {
	err := dr.db.Save(&dorm).Error

	return dorm, err
}

func (dr *dormRepository) DeleteDorm(dorm models.Dorm) (models.Dorm, error) {
	err := dr.db.Delete(&dorm).Error

	return dorm, err
}
