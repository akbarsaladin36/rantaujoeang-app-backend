package services

import (
	"rantaujoeang-app-backend/helpers"
	"rantaujoeang-app-backend/inputs"
	"rantaujoeang-app-backend/models"
	"rantaujoeang-app-backend/repositories"
	"time"
)

type PaymentService interface {
	FindPaymentsService() ([]models.Payment, error)
	FindPaymentsByUserIdService(user_uuid string) ([]models.Payment, error)
	FindPaymentService(payment_code string) (models.Payment, error)
	CreatePaymentService(paymentInput inputs.PaymentInput, currentUser map[string]string) (models.Payment, error)
	UpdatePaymentService(payment_code string, paymentStatus string) error
	DeletePaymentService(payment_code string) (models.Payment, error)
}

type paymentService struct {
	paymentRepository repositories.PaymentRepository
}

func NewPaymentService(paymentRepository repositories.PaymentRepository) *paymentService {
	return &paymentService{paymentRepository}
}

func (ps *paymentService) FindPaymentsService() ([]models.Payment, error) {
	payments, err := ps.paymentRepository.FindPayments()

	return payments, err
}

func (ps *paymentService) FindPaymentsByUserIdService(user_uuid string) ([]models.Payment, error) {
	payments, err := ps.paymentRepository.FindPaymentsByUserId(user_uuid)

	return payments, err
}

func (ps *paymentService) FindPaymentService(payment_code string) (models.Payment, error) {
	payment, err := ps.paymentRepository.FindPayment(payment_code)

	return payment, err
}

func (ps *paymentService) CreatePaymentService(paymentInput inputs.PaymentInput, currentUser map[string]string) (models.Payment, error) {
	paymentCode := helpers.GenerateCode("pt")

	payment := models.Payment{
		DormCode:                   paymentInput.DormCode,
		UserUuid:                   currentUser["user_uuid"],
		PaymentCode:                paymentCode,
		PaymentAmount:              paymentInput.PaymentAmount,
		PaymentQuantityAmount:      paymentInput.PaymentQuantityAmount,
		PaymentDescription:         paymentInput.PaymentDescription,
		PaymentTags:                paymentInput.PaymentTags,
		PaymentStatusCd:            "pending",
		PaymentCreatedDate:         time.Now(),
		PaymentCreatedUserUuid:     currentUser["user_uuid"],
		PaymentCreatedUserUsername: currentUser["user_username"],
	}

	newPayment, err := ps.paymentRepository.CreatePayment(payment)

	return newPayment, err
}

func (ps *paymentService) UpdatePaymentService(payment_code string, paymentStatus string) error {
	err := ps.paymentRepository.UpdatePayment(payment_code, paymentStatus)

	return err
}

func (ps *paymentService) DeletePaymentService(payment_code string) (models.Payment, error) {
	checkPayment, _ := ps.paymentRepository.FindPayment(payment_code)

	deletePayment, err := ps.paymentRepository.DeletePayment(checkPayment)

	return deletePayment, err
}
