package responses

import (
	"rantaujoeang-app-backend/models"
	"time"
)

type InvoiceResponse struct {
	InvoiceCode           string `json:"invoice_code"`
	PaymentCode           string `json:"payment_code"`
	DormCode              string `json:"dorm_code"`
	InvoiceAmount         string `json:"invoice_amount"`
	InvoiceQuantityAmount string `json:"invoice_quantity_amount"`
	InvoiceDescription    string `json:"invoice_description"`
	InvoiceTags           string `json:"invoice_tags"`
	InvoiceStatusCd       string `json:"invoice_status_cd"`
}

type CreateInvoiceResponse struct {
	InvoiceCode                string    `json:"invoice_code"`
	PaymentCode                string    `json:"payment_code"`
	DormCode                   string    `json:"dorm_code"`
	InvoiceAmount              string    `json:"invoice_amount"`
	InvoiceQuantityAmount      string    `json:"invoice_quantity_amount"`
	InvoiceDescription         string    `json:"invoice_description"`
	InvoiceTags                string    `json:"invoice_tags"`
	InvoiceStatusCd            string    `json:"invoice_status_cd"`
	InvoiceCreatedDate         time.Time `json:"invoice_created_date"`
	InvoiceCreatedUserUuid     string    `json:"invoice_created_user_uuid"`
	InvoiceCreatedUserUsername string    `json:"invoice_created_user_username"`
}

type UpdateInvoiceResponse struct {
	InvoiceCode                string    `json:"invoice_code"`
	PaymentCode                string    `json:"payment_code"`
	DormCode                   string    `json:"dorm_code"`
	InvoiceAmount              string    `json:"invoice_amount"`
	InvoiceQuantityAmount      string    `json:"invoice_quantity_amount"`
	InvoiceDescription         string    `json:"invoice_description"`
	InvoiceTags                string    `json:"invoice_tags"`
	InvoiceStatusCd            string    `json:"invoice_status_cd"`
	InvoiceUpdatedDate         time.Time `json:"invoice_updated_date"`
	InvoiceUpdatedUserUuid     string    `json:"invoice_updated_user_uuid"`
	InvoiceUpdatedUserUsername string    `json:"invoice_updated_user_username"`
}

func GetInvoiceResponse(invoiceRsps models.Invoice) InvoiceResponse {
	return InvoiceResponse{
		InvoiceCode:           invoiceRsps.InvoiceCode,
		PaymentCode:           invoiceRsps.PaymentCode,
		DormCode:              invoiceRsps.DormCode,
		InvoiceAmount:         invoiceRsps.InvoiceAmount,
		InvoiceQuantityAmount: invoiceRsps.InvoiceQuantityAmount,
		InvoiceDescription:    invoiceRsps.InvoiceDescription,
		InvoiceTags:           invoiceRsps.InvoiceTags,
		InvoiceStatusCd:       invoiceRsps.InvoiceStatusCd,
	}
}

func GetCreateInvoiceResponse(invoiceRsps models.Invoice) CreateInvoiceResponse {
	return CreateInvoiceResponse{
		InvoiceCode:                invoiceRsps.InvoiceCode,
		PaymentCode:                invoiceRsps.PaymentCode,
		DormCode:                   invoiceRsps.DormCode,
		InvoiceAmount:              invoiceRsps.InvoiceAmount,
		InvoiceQuantityAmount:      invoiceRsps.InvoiceQuantityAmount,
		InvoiceDescription:         invoiceRsps.InvoiceDescription,
		InvoiceTags:                invoiceRsps.InvoiceTags,
		InvoiceStatusCd:            invoiceRsps.InvoiceStatusCd,
		InvoiceCreatedDate:         invoiceRsps.InvoiceCreatedDate,
		InvoiceCreatedUserUuid:     invoiceRsps.InvoiceCreatedUserUuid,
		InvoiceCreatedUserUsername: invoiceRsps.InvoiceCreatedUserUsername,
	}
}

func GetUpdateInvoiceResponse(invoiceRsps models.Invoice) UpdateInvoiceResponse {
	return UpdateInvoiceResponse{
		InvoiceCode:                invoiceRsps.InvoiceCode,
		PaymentCode:                invoiceRsps.PaymentCode,
		DormCode:                   invoiceRsps.DormCode,
		InvoiceAmount:              invoiceRsps.InvoiceAmount,
		InvoiceQuantityAmount:      invoiceRsps.InvoiceQuantityAmount,
		InvoiceDescription:         invoiceRsps.InvoiceDescription,
		InvoiceTags:                invoiceRsps.InvoiceTags,
		InvoiceStatusCd:            invoiceRsps.InvoiceStatusCd,
		InvoiceUpdatedDate:         invoiceRsps.InvoiceUpdatedDate,
		InvoiceUpdatedUserUuid:     invoiceRsps.InvoiceUpdatedUserUuid,
		InvoiceUpdatedUserUsername: invoiceRsps.InvoiceUpdatedUserUsername,
	}
}
