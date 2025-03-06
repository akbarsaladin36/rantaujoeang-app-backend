package inputs

type PaymentInput struct {
	DormCode              string `json:"dorm_code"`
	UserUuid              string `json:"user_uuid"`
	PaymentCode           string `json:"payment_code"`
	PaymentAmount         string `json:"payment_amount"`
	PaymentQuantityAmount string `json:"payment_quantity_amount"`
	PaymentDescription    string `json:"payment_description"`
	PaymentTags           string `json:"payment_tags"`
	PaymentStatusCd       string `json:"payment_status_cd"`
}
