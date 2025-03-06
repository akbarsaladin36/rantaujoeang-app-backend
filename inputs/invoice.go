package inputs

type CreateInvoiceInput struct {
	PaymentCode           string `json:"payment_code"`
	DormCode              string `json:"dorm_code"`
	InvoiceAmount         string `json:"invoice_amount"`
	InvoiceQuantityAmount string `json:"invoice_quantity_amount"`
	InvoiceDescription    string `json:"invoice_description"`
	InvoiceTags           string `json:"invoice_tags"`
	InvoiceStatusCd       string `json:"invoice_status_cd"`
}

type UpdateInvoiceInput struct {
	InvoiceStatusCd string `json:"invoice_status_cd"`
}
