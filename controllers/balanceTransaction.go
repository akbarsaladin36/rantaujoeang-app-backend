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

type balanceTransactionController struct {
	balanceTransactionService services.BalanceTransactionService
}

func NewBalanceTransactionController(balanceTransactionService services.BalanceTransactionService) *balanceTransactionController {
	return &balanceTransactionController{balanceTransactionService}
}

func (btc *balanceTransactionController) FindBalanceTransactionsController(c *gin.Context) {
	balanceTransactions, err := btc.balanceTransactionService.FindBalanceTransactionsService()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "All balance transactions data is empty! Please create a new balance transaction!",
		})
		return
	}

	var balanceTransactionsRsps []responses.BalanceTransactionRsps

	for _, balancebalanceTransaction := range balanceTransactions {
		balanceTransactionRsps := responses.GetBalanceTransactionResponse(balancebalanceTransaction)

		balanceTransactionsRsps = append(balanceTransactionsRsps, balanceTransactionRsps)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All balance transactions data is succesfully appeared!",
		"data":    balanceTransactionsRsps,
	})
}

func (btc *balanceTransactionController) FindBalanceTransactionsByUserIdController(c *gin.Context) {
	currentUser := middleware.CurrentUser(c)

	user_uuid := currentUser["user_uuid"]

	balanceTransactions, err := btc.balanceTransactionService.FindBalanceTransactionByUserIdService(user_uuid)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "All balance transactions data is empty! Please create a new balance transaction!",
		})
		return
	}

	var balanceTransactionsRsps []responses.BalanceTransactionRsps

	for _, balancebalanceTransaction := range balanceTransactions {
		balanceTransactionRsps := responses.GetBalanceTransactionResponse(balancebalanceTransaction)

		balanceTransactionsRsps = append(balanceTransactionsRsps, balanceTransactionRsps)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All balance transactions data based current user is succesfully appeared!",
		"data":    balanceTransactionsRsps,
	})
}

func (btc *balanceTransactionController) FindBalanceTransactionController(c *gin.Context) {
	balance_transaction_code := c.Param("balance_transaction_code")

	balanceTransaction, err := btc.balanceTransactionService.FindBalanceTransactionService(balance_transaction_code)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "400",
			"message": "Your balance code " + balance_transaction_code + " is not found!",
		})
		return
	}

	balanceTransactionRsps := responses.GetBalanceTransactionResponse(balanceTransaction)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Your balance code " + balance_transaction_code + " is succesfully appeared!",
		"data":    balanceTransactionRsps,
	})
}

func (btc *balanceTransactionController) CreateBalanceTransactionController(c *gin.Context) {
	var createBalanceTransationInput inputs.CreateBalanceTransactionInput

	errCreateBalanceTransaction := c.ShouldBindJSON(&createBalanceTransationInput)

	if errCreateBalanceTransaction != nil {
		errorMessages := []string{}
		for _, e := range errCreateBalanceTransaction.(validator.ValidationErrors) {
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

	newBalanceTransaction, errNewBalanceTransaction := btc.balanceTransactionService.CreateBalanceTransactionService(createBalanceTransationInput, currentUser)

	if errNewBalanceTransaction != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "401",
			"message": "Something wrong when processing create a new balance transaction! Please try again!",
		})
		return
	}

	newBalanceTransactionRsps := responses.GetBalanceTransactionResponse(newBalanceTransaction)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A new balance transaction is succesfully created!",
		"data":    newBalanceTransactionRsps,
	})

}
