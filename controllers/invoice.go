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

type invoiceController struct {
	invoiceService services.InvoiceService
}

func NewInvoiceController(invoiceService services.InvoiceService) *invoiceController {
	return &invoiceController{invoiceService}
}

func (ic *invoiceController) FindInvoicesController(c *gin.Context) {
	invoices, err := ic.invoiceService.FindInvoicesService()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "All invoice data are empty! Please create a new invoice now!",
		})
		return
	}

	var invoicesRsps []responses.InvoiceResponse

	for _, invoice := range invoices {
		invoiceRsps := responses.GetInvoiceResponse(invoice)

		invoicesRsps = append(invoicesRsps, invoiceRsps)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All invoices are succesfully appeared!",
		"data":    invoicesRsps,
	})
}

func (ic *invoiceController) FindInvoicesByUserIdController(c *gin.Context) {
	currentUser := middleware.CurrentUser(c)

	invoices, err := ic.invoiceService.FindInvoicesByUserIdService(currentUser["user_uuid"])

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "All invoice data are empty! Please create a new invoice now!",
		})
		return
	}

	var invoicesRsps []responses.InvoiceResponse

	for _, invoice := range invoices {
		invoiceRsps := responses.GetInvoiceResponse(invoice)

		invoicesRsps = append(invoicesRsps, invoiceRsps)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All invoices are succesfully appeared!",
		"data":    invoicesRsps,
	})
}

func (ic *invoiceController) FindInvoiceController(c *gin.Context) {
	invoice_code := c.Param("invoice_code")

	invoice, err := ic.invoiceService.FindInvoiceService(invoice_code)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "400",
			"message": "An invoice data " + invoice_code + " is not found! Please try find a different invoice code again!",
		})
		return
	}

	invoiceRsps := responses.GetInvoiceResponse(invoice)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "An invoice data " + invoice_code + " is succesfully appeared!",
		"data":    invoiceRsps,
	})
}

func (ic *invoiceController) CreateInvoiceController(c *gin.Context) {
	var createInvoiceInput inputs.CreateInvoiceInput

	errCreateInvoiceInput := c.ShouldBindJSON(&createInvoiceInput)

	if errCreateInvoiceInput != nil {
		errorMessages := []string{}
		for _, e := range errCreateInvoiceInput.(validator.ValidationErrors) {
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

	newInvoice, err := ic.invoiceService.CreateInvoiceService(createInvoiceInput, currentUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Error when creating a new invoice! Please try again!",
		})
		return
	}

	createInvoiceRsps := responses.GetCreateInvoiceResponse(newInvoice)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A new invoice is succesfully created!",
		"data":    createInvoiceRsps,
	})
}

func (ic *invoiceController) UpdateInvoiceController(c *gin.Context) {
	invoice_code := c.Param("invoice_code")

	_, err := ic.invoiceService.FindInvoiceService(invoice_code)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "400",
			"message": "An invoice data " + invoice_code + " is not found! Please try find a different invoice code again!",
		})
		return
	}

	var updateInvoiceInput inputs.UpdateInvoiceInput

	errUpdateInvoiceInput := c.ShouldBindJSON(&updateInvoiceInput)

	if errUpdateInvoiceInput != nil {
		errorMessages := []string{}
		for _, e := range errUpdateInvoiceInput.(validator.ValidationErrors) {
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

	updateInvoice, errUpdateInvoice := ic.invoiceService.UpdateInvoiceService(invoice_code, updateInvoiceInput, currentUser)

	if errUpdateInvoice != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Error when updating an existing invoice! Please try again!",
		})
		return
	}

	updateInvoiceRsps := responses.GetUpdateInvoiceResponse(updateInvoice)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "An invoice data " + invoice_code + " is succesfully updated!",
		"data":    updateInvoiceRsps,
	})
}

func (ic *invoiceController) DeleteInvoiceController(c *gin.Context) {
	invoice_code := c.Param("invoice_code")

	_, err := ic.invoiceService.FindInvoiceService(invoice_code)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "400",
			"message": "An invoice data " + invoice_code + " is not found! Please try find a different invoice code again!",
		})
		return
	}

	_, errDeleteInvoice := ic.invoiceService.DeleteInvoiceService(invoice_code)

	if errDeleteInvoice != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Error when deleting an existing invoice! Please try again!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "An invoice data " + invoice_code + " is succesfully deleted!",
	})
}
