package controllers

import (
	"fmt"
	"net/http"
	"rantaujoeang-app-backend/helpers"
	"rantaujoeang-app-backend/inputs"
	"rantaujoeang-app-backend/middleware"
	"rantaujoeang-app-backend/models"
	"rantaujoeang-app-backend/responses"
	"rantaujoeang-app-backend/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *authController {
	return &authController{authService}
}

func (ac *authController) RegisterController(c *gin.Context) {
	var registerInput inputs.RegisterInput

	errRegisterInput := c.ShouldBindJSON(&registerInput)

	if errRegisterInput != nil {
		errorMessages := []string{}
		for _, e := range errRegisterInput.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})
		return
	}

	_, errCheckUser := ac.authService.FindUserService(registerInput.Username)

	if errCheckUser == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "Username is exist! Please register with another username now!",
		})
		return
	}

	newUser, errNewUser := ac.authService.RegisterService(registerInput)

	if errNewUser != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "Process registering a new user is not finished! Please try again!",
		})
		return
	}

	registerRsps := responses.GetRegisterResponse(newUser)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A new user is succesfully registered!",
		"data":    registerRsps,
	})
}

func (ac *authController) LoginController(c *gin.Context) {
	var loginInput inputs.LoginInput

	errLoginInput := c.ShouldBindJSON(&loginInput)

	if errLoginInput != nil {
		errorMessages := []string{}
		for _, e := range errLoginInput.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})
		return
	}

	_, errCheckUser := ac.authService.FindUserService(loginInput.Username)

	if errCheckUser != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "Username is not exist! Please register now!",
		})
		return
	}

	loginUser, errLoginUser := ac.authService.LoginService(loginInput)

	if errLoginUser != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "Error when processing login user! Please try again!",
		})
		return
	}

	checkSessionUser, errCheckSessionUser := middleware.CheckSessionUser(loginUser.UserUUID)

	tokenString := helpers.GenerateSessionToken(300)

	if errCheckSessionUser != nil {

		session := models.Session{
			SessionToken:        tokenString,
			SessionUserUUID:     loginUser.UserUUID,
			SessionUserUsername: loginUser.UserUsername,
			SessionUserRole:     loginUser.UserRole,
			SessionStartAt:      time.Now(),
			SessionExpiredAt:    time.Now().Add(24 * time.Hour),
			SessionCreatedAt:    time.Now(),
		}

		sessionUser, _ := middleware.CreateSessionUser(session)

		loginUserRsps := responses.GetLoginResponse(sessionUser, tokenString)

		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "A user is succesfully login!",
			"data":    loginUserRsps,
		})
	} else {
		checkSessionUser.SessionToken = tokenString
		checkSessionUser.SessionUserUUID = loginUser.UserUUID
		checkSessionUser.SessionUserUsername = loginUser.UserUsername
		checkSessionUser.SessionUserRole = loginUser.UserRole
		checkSessionUser.SessionStartAt = time.Now()
		checkSessionUser.SessionExpiredAt = time.Now().Add(24 * time.Hour)
		checkSessionUser.SessionUpdateAt = time.Now()

		updateSessionUser, _ := middleware.UpdateSessionUser(checkSessionUser)

		loginUserRsps := responses.GetLoginResponse(updateSessionUser, tokenString)

		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "A user is succesfully login!",
			"data":    loginUserRsps,
		})
	}

}
