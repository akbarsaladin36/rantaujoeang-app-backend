package responses

import (
	"rantaujoeang-app-backend/models"
	"time"
)

type PaymentResponse struct {
	DormCode              string `json:"dorm_code"`
	UserUuid              string `json:"user_uuid"`
	PaymentCode           string `json:"payment_code"`
	PaymentAmount         string `json:"payment_amount"`
	PaymentQuantityAmount string `json:"payment_quantity_amount"`
	PaymentDescription    string `json:"payment_description"`
	PaymentTags           string `json:"payment_tags"`
	PaymentStatusCd       string `json:"payment_status_cd"`
}

type CreatePaymentResponse struct {
	DormCode                   string    `json:"dorm_code"`
	UserUuid                   string    `json:"user_uuid"`
	PaymentCode                string    `json:"payment_code"`
	PaymentAmount              string    `json:"payment_amount"`
	PaymentQuantityAmount      string    `json:"payment_quantity_amount"`
	PaymentDescription         string    `json:"payment_description"`
	PaymentTags                string    `json:"payment_tags"`
	PaymentStatusCd            string    `json:"payment_status_cd"`
	PaymentCreatedDate         time.Time `json:"payment_created_date"`
	PaymentCreatedUserUuid     string    `json:"payment_created_user_uuid"`
	PaymentCreatedUserUsername string    `json:"payment_created_user_username"`
}

func GetPaymentResponse(paymentRsps models.Payment) PaymentResponse {
	return PaymentResponse{
		DormCode:           paymentRsps.DormCode,
		UserUuid:           paymentRsps.UserUuid,
		PaymentCode:        paymentRsps.PaymentCode,
		PaymentAmount:      paymentRsps.PaymentAmount,
		PaymentDescription: paymentRsps.PaymentDescription,
		PaymentTags:        paymentRsps.PaymentTags,
		PaymentStatusCd:    paymentRsps.PaymentStatusCd,
	}
}

func GetCreatePaymentResponse(paymentRsps models.Payment) CreatePaymentResponse {
	return CreatePaymentResponse{
		DormCode:                   paymentRsps.DormCode,
		UserUuid:                   paymentRsps.UserUuid,
		PaymentCode:                paymentRsps.PaymentCode,
		PaymentAmount:              paymentRsps.PaymentAmount,
		PaymentDescription:         paymentRsps.PaymentDescription,
		PaymentTags:                paymentRsps.PaymentTags,
		PaymentStatusCd:            paymentRsps.PaymentStatusCd,
		PaymentCreatedDate:         paymentRsps.PaymentCreatedDate,
		PaymentCreatedUserUuid:     paymentRsps.PaymentCreatedUserUuid,
		PaymentCreatedUserUsername: paymentRsps.PaymentCreatedUserUsername,
	}
}
