package responses

import (
	"rantaujoeang-app-backend/models"
	"time"
)

type GetUserRsps struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}

type CreateUserRsps struct {
	Email           string    `json:"email"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	CreatedDate     time.Time `json:"created_user_date"`
	CreatedUserUuid string    `json:"created_user_uuid"`
	CreatedUsername string    `json:"created_username"`
}

type UpdateUserRsps struct {
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Address         string    `json:"address"`
	UpdatedDate     time.Time `json:"updates_user_date"`
	UpdatedUserUuid string    `json:"updated_user_uuid"`
	UpdatedUsername string    `json:"updated_username"`
}

func GetUserResponse(userRsps models.User) GetUserRsps {
	return GetUserRsps{
		Email:     userRsps.UserEmail,
		Username:  userRsps.UserUsername,
		FirstName: userRsps.UserFirstName,
		LastName:  userRsps.UserLastName,
		Address:   userRsps.UserAddress,
	}
}

func GetCreateUserRsps(userRsps models.User) CreateUserRsps {
	return CreateUserRsps{
		Email:           userRsps.UserEmail,
		Username:        userRsps.UserUsername,
		Password:        userRsps.UserPassword,
		CreatedDate:     userRsps.UserCreatedDate,
		CreatedUserUuid: userRsps.UserCreatedUserUuid,
		CreatedUsername: userRsps.UserCreatedUsername,
	}
}

func GetUpdateUserRsps(userRsps models.User) UpdateUserRsps {
	return UpdateUserRsps{
		FirstName:       userRsps.UserFirstName,
		LastName:        userRsps.UserLastName,
		Address:         userRsps.UserAddress,
		UpdatedDate:     userRsps.UserUpdatedDate,
		UpdatedUserUuid: userRsps.UserUpdatedUserUuid,
		UpdatedUsername: userRsps.UserUpdatedUsername,
	}
}
