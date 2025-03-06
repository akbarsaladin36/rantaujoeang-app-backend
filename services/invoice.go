package services

import (
	"rantaujoeang-app-backend/helpers"
	"rantaujoeang-app-backend/inputs"
	"rantaujoeang-app-backend/models"
	"rantaujoeang-app-backend/repositories"
	"strconv"
	"time"
)

type InvoiceService interface {
	FindInvoicesService() ([]models.Invoice, error)
	FindInvoicesByUserIdService(user_uuid string) ([]models.Invoice, error)
	FindInvoiceService(invoice_code string) (models.Invoice, error)
	CreateInvoiceService(createInvoiceInput inputs.CreateInvoiceInput, currentUser map[string]string) (models.Invoice, error)
	UpdateInvoiceService(invoice_code string, updateInvoiceInput inputs.UpdateInvoiceInput, currentUser map[string]string) (models.Invoice, error)
	DeleteInvoiceService(invoice_code string) (models.Invoice, error)
}

type invoiceService struct {
	invoiceRepository repositories.InvoiceRepository
}

func NewInvoiceService(invoiceRepository repositories.InvoiceRepository) *invoiceService {
	return &invoiceService{invoiceRepository}
}

func (is *invoiceService) FindInvoicesService() ([]models.Invoice, error) {
	invoices, err := is.invoiceRepository.FindInvoices()

	return invoices, err
}

func (is *invoiceService) FindInvoicesByUserIdService(user_uuid string) ([]models.Invoice, error) {
	invoices, err := is.invoiceRepository.FindInvoicesByUserId(user_uuid)

	return invoices, err
}

func (is *invoiceService) FindInvoiceService(invoice_code string) (models.Invoice, error) {
	invoice, err := is.invoiceRepository.FindInvoice(invoice_code)

	return invoice, err
}

func (is *invoiceService) CreateInvoiceService(createInvoiceInput inputs.CreateInvoiceInput, currentUser map[string]string) (models.Invoice, error) {
	invoiceCode := helpers.GenerateCode("inv")

	invoice := models.Invoice{
		PaymentCode:                createInvoiceInput.PaymentCode,
		DormCode:                   createInvoiceInput.DormCode,
		UserUuid:                   currentUser["user_uuid"],
		InvoiceCode:                invoiceCode,
		InvoiceAmount:              createInvoiceInput.InvoiceAmount,
		InvoiceQuantityAmount:      createInvoiceInput.InvoiceQuantityAmount,
		InvoiceDescription:         createInvoiceInput.InvoiceDescription,
		InvoiceTags:                createInvoiceInput.InvoiceTags,
		InvoiceStatusCd:            "pending",
		InvoiceCreatedDate:         time.Now(),
		InvoiceCreatedUserUuid:     currentUser["user_uuid"],
		InvoiceCreatedUserUsername: currentUser["user_username"],
	}

	newInvoice, err := is.invoiceRepository.CreateInvoice(invoice)

	return newInvoice, err
}

func (is *invoiceService) UpdateInvoiceService(invoice_code string, updateInvoiceInput inputs.UpdateInvoiceInput, currentUser map[string]string) (models.Invoice, error) {
	current_user_uuid := currentUser["user_uuid"]

	checkInvoice, _ := is.invoiceRepository.FindInvoice(invoice_code)

	switch updateInvoiceInput.InvoiceStatusCd {
	case "pending":
		checkInvoice.InvoiceStatusCd = updateInvoiceInput.InvoiceStatusCd
		checkInvoice.InvoiceUpdatedDate = time.Now()
		checkInvoice.InvoiceUpdatedUserUuid = currentUser["user_uuid"]
		checkInvoice.InvoiceUpdatedUserUsername = currentUser["user_username"]

		updateInvoice, err := is.invoiceRepository.UpdateInvoice(checkInvoice)

		return updateInvoice, err
	case "paid":
		checkInvoice.InvoiceStatusCd = updateInvoiceInput.InvoiceStatusCd
		checkInvoice.InvoiceUpdatedDate = time.Now()
		checkInvoice.InvoiceUpdatedUserUuid = currentUser["user_uuid"]
		checkInvoice.InvoiceUpdatedUserUsername = currentUser["user_username"]

		checkDorm, _ := is.invoiceRepository.FindDorm(checkInvoice.DormCode)
		checkUser, _ := is.invoiceRepository.FindUser(current_user_uuid)

		dormQuantity, _ := strconv.ParseInt(checkDorm.DormQuantityAmount, 0, 64)
		invoiceQuantity, _ := strconv.ParseInt(checkInvoice.InvoiceQuantityAmount, 0, 64)
		newResultQuantity := dormQuantity - invoiceQuantity
		newResultQuantityToString := strconv.FormatInt(newResultQuantity, 10)

		invoiceAmount, _ := strconv.ParseInt(checkInvoice.InvoiceAmount, 0, 64)
		newInvoiceAmount := invoiceQuantity * invoiceAmount
		userBalanceTransactionAmount, _ := strconv.ParseInt(checkUser.UserBalanceTransactionAmount, 0, 64)
		newUserBalanceTransactionAmount := userBalanceTransactionAmount - newInvoiceAmount
		newUserBalanceTransactionAmountToString := strconv.FormatInt(newUserBalanceTransactionAmount, 10)

		is.invoiceRepository.UpdateDormQuantity(checkInvoice.DormCode, newResultQuantityToString)
		is.invoiceRepository.UpdateUserBalanceTransaction(current_user_uuid, newUserBalanceTransactionAmountToString)
		is.invoiceRepository.UpdatePayment(checkInvoice.PaymentCode, "paid")

		updateInvoice, err := is.invoiceRepository.UpdateInvoice(checkInvoice)

		return updateInvoice, err
	default:
		checkInvoice.InvoiceStatusCd = updateInvoiceInput.InvoiceStatusCd
		checkInvoice.InvoiceUpdatedDate = time.Now()
		checkInvoice.InvoiceUpdatedUserUuid = currentUser["user_uuid"]
		checkInvoice.InvoiceUpdatedUserUsername = currentUser["user_username"]

		updateInvoice, err := is.invoiceRepository.UpdateInvoice(checkInvoice)

		return updateInvoice, err
	}
}

func (is *invoiceService) DeleteInvoiceService(invoice_code string) (models.Invoice, error) {
	checkInvoice, _ := is.invoiceRepository.FindInvoice(invoice_code)

	deleteInvoice, err := is.invoiceRepository.DeleteInvoice(checkInvoice)

	return deleteInvoice, err
}
