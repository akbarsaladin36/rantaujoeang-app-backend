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

type profileController struct {
	profileService services.ProfileService
}

func NewProfileController(profileService services.ProfileService) *profileController {
	return &profileController{profileService}
}

func (pc *profileController) FindProfileController(c *gin.Context) {
	currentUser := middleware.CurrentUser(c)

	username := currentUser["user_username"]

	checkProfile, errCheckProfile := pc.profileService.FindProfileService(username)

	if errCheckProfile != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "A username " + username + " data is not found!",
		})
		return
	}

	profileRsps := responses.GetUserResponse(checkProfile)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A username " + username + " data is succesfully appeared!",
		"data":    profileRsps,
	})
}

func (pc *profileController) UpdateProfileController(c *gin.Context) {
	currentUser := middleware.CurrentUser(c)

	username := currentUser["user_username"]

	_, errCheckProfile := pc.profileService.FindProfileService(username)

	if errCheckProfile != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "A username " + username + " data is not found!",
		})
		return
	}

	var updateProfileInput inputs.UpdateUserInput

	errUpdateProfileInput := c.ShouldBindJSON(&updateProfileInput)

	if errUpdateProfileInput != nil {
		errorMessages := []string{}
		for _, e := range errUpdateProfileInput.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})
		return
	}

	updateProfile, errUpdateProfile := pc.profileService.UpdateProfileService(updateProfileInput, currentUser)

	if errUpdateProfile != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "Something wrong when updating a profile for username " + username + "! Please try again!",
		})
		return
	}

	updateProfileRsps := responses.GetUpdateUserRsps(updateProfile)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A profile for username " + username + " data is succesfully updated!",
		"data":    updateProfileRsps,
	})
}
