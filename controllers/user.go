package controllers

import (
	"fmt"
	"net/http"
	"rantaujoeang-app-backend/inputs"
	"rantaujoeang-app-backend/middleware"
	"rantaujoeang-app-backend/responses"
	"rantaujoeang-app-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *userController {
	return &userController{userService}
}

func (uc *userController) FindUsersController(c *gin.Context) {
	users, err := uc.userService.FindUsersService()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "All users data is empty! Please create a new user!",
		})
		return
	}

	var usersRsps []responses.GetUserRsps

	for _, user := range users {
		userRsps := responses.GetUserResponse(user)

		usersRsps = append(usersRsps, userRsps)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All users data is succesfully appeared!",
		"data":    usersRsps,
	})
}

func (uc *userController) FindUserController(c *gin.Context) {
	username := c.Param("username")

	user, err := uc.userService.FindUserService(username)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "A username " + username + " data is not found!",
		})
		return
	}

	userRsps := responses.GetUserResponse(user)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A username " + username + " data is succesfully appeared!",
		"data":    userRsps,
	})
}

func (uc *userController) CreateUserController(c *gin.Context) {
	var createUserInput inputs.CreateUserInput

	errCreateUserInput := c.ShouldBindJSON(&createUserInput)

	if errCreateUserInput != nil {
		errorMessages := []string{}
		for _, e := range errCreateUserInput.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})
		return
	}

	_, errCheckUser := uc.userService.FindUserService(createUserInput.Username)

	if errCheckUser == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "A username " + createUserInput.Username + " data is exist! Please try to find a new username again!",
		})
		return
	}

	currentUser := middleware.CurrentUser(c)

	newUser, errNewUser := uc.userService.CreateUserService(createUserInput, currentUser)

	if errNewUser != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "Something wrong when processing create a new user! Please try again!",
		})
		return
	}

	newUserRsps := responses.GetCreateUserRsps(newUser)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A new user is succesfully registered!",
		"data":    newUserRsps,
	})
}

func (uc *userController) UpdateUserController(c *gin.Context) {
	username := c.Param("username")

	_, errCheckUser := uc.userService.FindUserService(username)

	if errCheckUser != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "A username " + username + " data is not found!",
		})
		return
	}

	var updateUserInput inputs.UpdateUserInput

	errUpdateUserInput := c.ShouldBindJSON(&updateUserInput)

	if errUpdateUserInput != nil {
		errorMessages := []string{}
		for _, e := range errUpdateUserInput.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})
		return
	}

	currentUser := middleware.CurrentUser(c)

	updateUser, errUpdateUser := uc.userService.UpdateUserService(username, updateUserInput, currentUser)

	if errUpdateUser != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "Something wrong when updating a data for username " + username + "! Please try again!",
		})
		return
	}

	updateUserRsps := responses.GetUpdateUserRsps(updateUser)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A username " + username + " data is succesfully updated!",
		"data":    updateUserRsps,
	})
}

func (uc *userController) DeleteUserController(c *gin.Context) {
	username := c.Param("username")

	_, errCheckUser := uc.userService.FindUserService(username)

	if errCheckUser != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "A username " + username + " data is not found!",
		})
		return
	}

	_, errDeleteUser := uc.userService.DeleteUserService(username)

	if errDeleteUser != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "Something wrong when deleting a data for username " + username + "! Please try again!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A username " + username + " data is succesfully updated!",
	})

}
