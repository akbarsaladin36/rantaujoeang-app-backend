package controllers

import (
	"fmt"
	"net/http"
	"rantaujoeang-app-backend/helpers"
	"rantaujoeang-app-backend/inputs"
	"rantaujoeang-app-backend/middleware"
	"rantaujoeang-app-backend/responses"
	"rantaujoeang-app-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type dormController struct {
	dormService services.DormService
}

func NewDormController(dormService services.DormService) *dormController {
	return &dormController{dormService}
}

func (dc *dormController) FindDormsController(c *gin.Context) {
	dorms, errDorms := dc.dormService.FindDormsService()

	if errDorms != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "All dorms data is empty! Please try create a new dorm now!",
			"status":  "200",
		})
		return
	}

	var dormsRsps []responses.DormResponse

	for _, dorm := range dorms {
		dormRsps := responses.GetDormResponse(dorm)

		dormsRsps = append(dormsRsps, dormRsps)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "All dorms data is succesfully appeared!",
		"status":  "200",
		"data":    dormsRsps,
	})
}

func (dc *dormController) FindDormController(c *gin.Context) {
	dorm_code := c.Param("dorm_code")

	dorm, err := dc.dormService.FindDormService(dorm_code)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "The dorm with code " + dorm_code + " is not found!",
			"status":  "400",
		})
		return
	}

	dormRsps := responses.GetDormResponse(dorm)

	c.JSON(http.StatusOK, gin.H{
		"message": "The dorm with code " + dorm_code + " is succesfully appeared!",
		"status":  "200",
		"data":    dormRsps,
	})
}

func (dc *dormController) CreateDormController(c *gin.Context) {
	var createDormInput inputs.DormInput

	errCreateDormInput := c.ShouldBindJSON(&createDormInput)

	if errCreateDormInput != nil {
		errorMessages := []string{}
		for _, e := range errCreateDormInput.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})
		return
	}

	dormCode := helpers.GenerateSlug(createDormInput.DormName)

	_, errCheckDorm := dc.dormService.FindDormService(dormCode)

	if errCheckDorm == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "The dorm with code " + dormCode + " registered! Please try find a new dorm name now!",
			"status":  "400",
		})
		return
	}

	currentUser := middleware.CurrentUser(c)

	newDorm, errNewDorm := dc.dormService.CreateDormService(createDormInput, currentUser)

	if errNewDorm != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error when process creating a new dorm data! Please try again!",
			"status":  "400",
		})
		return
	}

	newDormRsps := responses.GetCreateDormResponse(newDorm)

	c.JSON(http.StatusOK, gin.H{
		"message": "A new dorm is succesfully created!",
		"status":  "200",
		"data":    newDormRsps,
	})
}

func (dc *dormController) UpdateDormController(c *gin.Context) {
	dorm_code := c.Param("dorm_code")

	_, errCheckDorm := dc.dormService.FindDormService(dorm_code)

	if errCheckDorm != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "The dorm with code " + dorm_code + " is not found!",
			"status":  "400",
		})
		return
	}

	var updateDormInput inputs.DormInput

	errUpdateDormInput := c.ShouldBindJSON(&updateDormInput)

	if errUpdateDormInput != nil {
		errorMessages := []string{}
		for _, e := range errUpdateDormInput.(validator.ValidationErrors) {
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

	updateDorm, errUpdateDorm := dc.dormService.UpdateDormService(dorm_code, updateDormInput, currentUser)

	if errUpdateDorm != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error when process updating a dorm data! Please try again!",
			"status":  "400",
		})
		return
	}

	updateDormRsps := responses.GetUpdateDormResponse(updateDorm)

	c.JSON(http.StatusOK, gin.H{
		"message": "The dorm with code " + dorm_code + " data is succesfully updated!",
		"status":  "200",
		"data":    updateDormRsps,
	})
}

func (dc *dormController) DeleteDormController(c *gin.Context) {
	dorm_code := c.Param("dorm_code")

	_, errCheckDorm := dc.dormService.FindDormService(dorm_code)

	if errCheckDorm != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "The dorm with code " + dorm_code + " is not found!",
			"status":  "400",
		})
		return
	}

	_, errDeleteDorm := dc.dormService.DeleteDormService(dorm_code)

	if errDeleteDorm != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error when process deleting a dorm data! Please try again!",
			"status":  "400",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "The dorm with code " + dorm_code + " data is succesfully deleted!",
		"status":  "200",
	})
}
