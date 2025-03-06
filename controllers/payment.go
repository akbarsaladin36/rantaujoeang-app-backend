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

type paymentController struct {
	paymentService services.PaymentService
}

func NewPaymentController(paymentService services.PaymentService) *paymentController {
	return &paymentController{paymentService}
}

func (pc *paymentController) FindPaymentsController(c *gin.Context) {
	payments, err := pc.paymentService.FindPaymentsService()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "All payments data are empty! Please create a new payment!",
		})
		return
	}

	var paymentsRsps []responses.PaymentResponse

	for _, payment := range payments {
		paymentRsps := responses.GetPaymentResponse(payment)

		paymentsRsps = append(paymentsRsps, paymentRsps)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All payments data are succesfully appeared!",
		"data":    paymentsRsps,
	})
}

func (pc *paymentController) FindPaymentsByUserIdController(c *gin.Context) {
	currentUser := middleware.CurrentUser(c)

	payments, err := pc.paymentService.FindPaymentsByUserIdService(currentUser["user_uuid"])

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "All payments data for user " + currentUser["user_username"] + " is empty!",
		})
		return
	}

	var paymentsRsps []responses.PaymentResponse

	for _, payment := range payments {
		paymentRsps := responses.GetPaymentResponse(payment)

		paymentsRsps = append(paymentsRsps, paymentRsps)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All payments data for user " + currentUser["user_username"] + " is succesfully appeared!",
		"data":    paymentsRsps,
	})
}

func (pc *paymentController) FindPaymentController(c *gin.Context) {
	payment_code := c.Param("payment_code")

	payment, err := pc.paymentService.FindPaymentService(payment_code)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "A payment code " + payment_code + " data are not found!",
		})
		return
	}

	paymentRsps := responses.GetPaymentResponse(payment)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A payment code " + payment_code + " data are succesfully appeared!",
		"data":    paymentRsps,
	})
}

func (pc *paymentController) CreatePaymentController(c *gin.Context) {
	var createPaymentInput inputs.PaymentInput

	errCreatePaymentInput := c.ShouldBindJSON(&createPaymentInput)

	if errCreatePaymentInput != nil {
		errorMessages := []string{}
		for _, e := range errCreatePaymentInput.(validator.ValidationErrors) {
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

	newPayment, err := pc.paymentService.CreatePaymentService(createPaymentInput, currentUser)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "Error when process creating new payment! Please try again!",
		})
		return
	}

	createPaymentRsps := responses.GetCreatePaymentResponse(newPayment)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A new payment is succesfully created!",
		"data":    createPaymentRsps,
	})
}

func (pc *paymentController) UpdatePaymentController(c *gin.Context) {
	paymentCode := c.Param("payment_code")

	_, errCheckPayment := pc.paymentService.FindPaymentService(paymentCode)

	if errCheckPayment != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "A payment code " + paymentCode + " data are not found!",
		})
		return
	}

	var updatePaymentInput inputs.PaymentInput

	paymentStatusCd := updatePaymentInput.PaymentStatusCd

	errUpdatePayment := pc.paymentService.UpdatePaymentService(paymentCode, paymentStatusCd)

	if errUpdatePayment != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "400",
			"message": "Error when updating status payment! Please try again!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A status for payment code " + paymentCode + " data are succesfully updated!",
	})
}

func (pc *paymentController) DeletePaymentController(c *gin.Context) {
	paymentCode := c.Param("payment_code")

	_, errCheckPayment := pc.paymentService.FindPaymentService(paymentCode)

	if errCheckPayment != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "A payment code " + paymentCode + " data are not found!",
		})
		return
	}

	_, errDeletePayment := pc.paymentService.DeletePaymentService(paymentCode)

	if errDeletePayment != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "400",
			"message": "Error when deleting payment data! Please try again!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A payment code " + paymentCode + " data are succesfully deleted!",
	})

}
