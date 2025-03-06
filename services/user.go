package services

import (
	"rantaujoeang-app-backend/helpers"
	"rantaujoeang-app-backend/inputs"
	"rantaujoeang-app-backend/models"
	"rantaujoeang-app-backend/repositories"
	"strings"
	"time"
)

type UserService interface {
	FindUsersService() ([]models.User, error)
	FindUserService(username string) (models.User, error)
	CreateUserService(createUserInput inputs.CreateUserInput, currentUser map[string]string) (models.User, error)
	UpdateUserService(username string, updateUserInput inputs.UpdateUserInput, currentUser map[string]string) (models.User, error)
	DeleteUserService(username string) (models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userService {
	return &userService{userRepository}
}

func (us *userService) FindUsersService() ([]models.User, error) {
	users, err := us.userRepository.FindUsers()

	return users, err
}

func (us *userService) FindUserService(username string) (models.User, error) {
	user, err := us.userRepository.FindUser(username)

	return user, err
}

func (us *userService) CreateUserService(createUserInput inputs.CreateUserInput, currentUser map[string]string) (models.User, error) {
	hashedPassword, _ := helpers.HashPassword(createUserInput.Password)

	convertUsernameToUUID := strings.ReplaceAll(helpers.GenerateUUID(createUserInput.Username), "-", "")

	user := models.User{
		UserUUID:            convertUsernameToUUID,
		UserEmail:           createUserInput.Email,
		UserUsername:        createUserInput.Username,
		UserPassword:        hashedPassword,
		UserRole:            "user",
		UserStatusCd:        "active",
		UserCreatedDate:     time.Now(),
		UserCreatedUserUuid: currentUser["user_uuid"],
		UserCreatedUsername: currentUser["user_username"],
	}

	newUser, err := us.userRepository.CreateUser(user)

	return newUser, err
}

func (us *userService) UpdateUserService(username string, updateUserInput inputs.UpdateUserInput, currentUser map[string]string) (models.User, error) {
	checkUser, _ := us.userRepository.FindUser(username)

	checkUser.UserFirstName = updateUserInput.FirstName
	checkUser.UserLastName = updateUserInput.LastName
	checkUser.UserAddress = updateUserInput.Address
	checkUser.UserUpdatedDate = time.Now()
	checkUser.UserUpdatedUserUuid = currentUser["user_uuid"]
	checkUser.UserUpdatedUsername = currentUser["user_username"]

	updateUser, err := us.userRepository.UpdateUser(checkUser)

	return updateUser, err
}

func (us *userService) DeleteUserService(username string) (models.User, error) {
	checkUser, _ := us.userRepository.FindUser(username)

	deleteUser, err := us.userRepository.DeleteUser(checkUser)

	return deleteUser, err
}
