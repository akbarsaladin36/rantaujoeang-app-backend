package services

import (
	"rantaujoeang-app-backend/helpers"
	"rantaujoeang-app-backend/inputs"
	"rantaujoeang-app-backend/models"
	"rantaujoeang-app-backend/repositories"
	"time"
)

type DormService interface {
	FindDormsService() ([]models.Dorm, error)
	FindDormService(dorm_code string) (models.Dorm, error)
	CreateDormService(dormInput inputs.DormInput, currentUser map[string]string) (models.Dorm, error)
	UpdateDormService(dorm_code string, dormInput inputs.DormInput, currentUser map[string]string) (models.Dorm, error)
	DeleteDormService(dorm_code string) (models.Dorm, error)
}

type dormService struct {
	dormRepository repositories.DormRepository
}

func NewDormService(dormRepository repositories.DormRepository) *dormService {
	return &dormService{dormRepository}
}

func (ds *dormService) FindDormsService() ([]models.Dorm, error) {
	dorms, err := ds.dormRepository.FindDorms()

	return dorms, err
}

func (ds *dormService) FindDormService(dorm_code string) (models.Dorm, error) {
	dorm, err := ds.dormRepository.FindDorm(dorm_code)

	return dorm, err
}

func (ds *dormService) CreateDormService(dormInput inputs.DormInput, currentUser map[string]string) (models.Dorm, error) {
	dormCode := helpers.GenerateSlug(dormInput.DormName)

	dorm := models.Dorm{
		DormCode:                dormCode,
		DormName:                dormInput.DormName,
		DormPriceAmount:         dormInput.DormPriceAmount,
		DormQuantityAmount:      dormInput.DormQuantityAmount,
		DormDescription:         dormInput.DormDescription,
		DormAddress:             dormInput.DormAddress,
		DormPhoneNumber:         dormInput.DormPhoneNumber,
		DormStatusCd:            "active",
		DormCreatedDate:         time.Now(),
		DormCreatedUserUuid:     currentUser["user_uuid"],
		DormCreatedUserUsername: currentUser["user_username"],
	}

	newDorm, err := ds.dormRepository.CreateDorm(dorm)

	return newDorm, err
}

func (ds *dormService) UpdateDormService(dorm_code string, dormInput inputs.DormInput, currentUser map[string]string) (models.Dorm, error) {
	dorm, _ := ds.dormRepository.FindDorm(dorm_code)

	dormCode := helpers.GenerateSlug(dormInput.DormName)

	dorm.DormCode = dormCode
	dorm.DormName = dormInput.DormName
	dorm.DormPriceAmount = dormInput.DormPriceAmount
	dorm.DormQuantityAmount = dormInput.DormQuantityAmount
	dorm.DormDescription = dormInput.DormDescription
	dorm.DormAddress = dormInput.DormAddress
	dorm.DormPhoneNumber = dormInput.DormPhoneNumber
	dorm.DormStatusCd = "active"
	dorm.DormUpdatedDate = time.Now()
	dorm.DormUpdatedUserUuid = currentUser["user_uuid"]
	dorm.DormUpdatedUserUsername = currentUser["user_username"]

	updatedDorm, err := ds.dormRepository.UpdateDorm(dorm)

	return updatedDorm, err
}

func (ds *dormService) DeleteDormService(dorm_code string) (models.Dorm, error) {
	dorm, _ := ds.dormRepository.FindDorm(dorm_code)

	deleteDorm, err := ds.dormRepository.DeleteDorm(dorm)

	return deleteDorm, err
}
