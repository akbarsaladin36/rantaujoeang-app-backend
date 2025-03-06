package responses

import "rantaujoeang-app-backend/models"

type RegisterRsps struct {
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}

type LoginRsps struct {
	UserUuid     string `json:"user_uuid"`
	UserUsername string `json:"user_username"`
	TokenString  string `json:"token_string"`
}

func GetRegisterResponse(userRsps models.User) RegisterRsps {
	return RegisterRsps{
		UserUsername: userRsps.UserUsername,
		UserEmail:    userRsps.UserEmail,
		UserPassword: userRsps.UserPassword,
	}
}

func GetLoginResponse(userRsps models.Session, tokenString string) LoginRsps {
	return LoginRsps{
		UserUuid:     userRsps.SessionUserUUID,
		UserUsername: userRsps.SessionUserUsername,
		TokenString:  tokenString,
	}
}
