package services

import (
	"rantaujoeang-app-backend/inputs"
	"rantaujoeang-app-backend/models"
	"rantaujoeang-app-backend/repositories"
	"time"
)

type ProfileService interface {
	FindProfileService(username string) (models.User, error)
	UpdateProfileService(updateUserInput inputs.UpdateUserInput, currentUser map[string]string) (models.User, error)
}

type profileService struct {
	profileRepository repositories.ProfileRepository
}

func NewProfileService(profileRepository repositories.ProfileRepository) *profileService {
	return &profileService{profileRepository}
}

func (ps *profileService) FindProfileService(username string) (models.User, error) {
	profile, err := ps.profileRepository.FindProfile(username)

	return profile, err
}

func (ps *profileService) UpdateProfileService(updateUserInput inputs.UpdateUserInput, currentUser map[string]string) (models.User, error) {
	user_uuid := currentUser["user_uuid"]
	username := currentUser["user_username"]

	profile, _ := ps.profileRepository.FindProfile(username)

	profile.UserFirstName = updateUserInput.FirstName
	profile.UserLastName = updateUserInput.LastName
	profile.UserAddress = updateUserInput.Address
	profile.UserUpdatedDate = time.Now()
	profile.UserUpdatedUserUuid = user_uuid
	profile.UserUpdatedUsername = username

	updateProfile, err := ps.profileRepository.UpdateProfile(profile)

	return updateProfile, err
}
