package services

import (
	"rantaujoeang-app-backend/helpers"
	"rantaujoeang-app-backend/inputs"
	"rantaujoeang-app-backend/models"
	"rantaujoeang-app-backend/repositories"
	"strings"
	"time"
)

type AuthService interface {
	FindUserService(username string) (models.User, error)
	RegisterService(registerInput inputs.RegisterInput) (models.User, error)
	LoginService(loginInput inputs.LoginInput) (models.User, error)
}

type authService struct {
	authRepository repositories.AuthRepository
}

func NewAuthService(authRepository repositories.AuthRepository) *authService {
	return &authService{authRepository}
}

func (as *authService) FindUserService(username string) (models.User, error) {
	user, err := as.authRepository.FindUser(username)

	return user, err
}

func (as *authService) RegisterService(registerInput inputs.RegisterInput) (models.User, error) {
	hashedPassword, _ := helpers.HashPassword(registerInput.Password)

	convertUsernameToUUID := strings.ReplaceAll(helpers.GenerateUUID(registerInput.Username), "-", "")

	user := models.User{
		UserUUID:            convertUsernameToUUID,
		UserEmail:           registerInput.Email,
		UserUsername:        registerInput.Username,
		UserPassword:        hashedPassword,
		UserRole:            "user",
		UserStatusCd:        "active",
		UserCreatedDate:     time.Now(),
		UserCreatedUserUuid: convertUsernameToUUID,
		UserCreatedUsername: registerInput.Username,
	}

	newUser, err := as.authRepository.CreateUser(user)

	return newUser, err
}

func (as *authService) LoginService(loginInput inputs.LoginInput) (models.User, error) {
	checkUser, err := as.authRepository.FindUser(loginInput.Username)

	helpers.CheckPassword(checkUser.UserPassword, loginInput.Password)

	return checkUser, err
}
