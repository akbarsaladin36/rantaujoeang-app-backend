package repositories

import (
	"rantaujoeang-app-backend/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	FindPayments() ([]models.Payment, error)
	FindPaymentsByUserId(user_uuid string) ([]models.Payment, error)
	FindPayment(payment_code string) (models.Payment, error)
	CreatePayment(payment models.Payment) (models.Payment, error)
	UpdatePayment(payment_code string, paymentStatus string) error
	DeletePayment(payment models.Payment) (models.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *paymentRepository {
	return &paymentRepository{db}
}

func (pr *paymentRepository) FindPayments() ([]models.Payment, error) {
	var payments []models.Payment

	err := pr.db.Find(&payments).Error

	return payments, err
}

func (pr *paymentRepository) FindPaymentsByUserId(user_uuid string) ([]models.Payment, error) {
	var payments []models.Payment

	err := pr.db.Where("user_uuid = ?", user_uuid).Find(&payments).Error

	return payments, err
}

func (pr *paymentRepository) FindPayment(payment_code string) (models.Payment, error) {
	var payment models.Payment

	err := pr.db.Where("payment_code = ?", payment_code).First(&payment).Error

	return payment, err
}

func (pr *paymentRepository) CreatePayment(payment models.Payment) (models.Payment, error) {
	err := pr.db.Create(&payment).Error

	return payment, err
}

func (pr *paymentRepository) UpdatePayment(payment_code string, paymentStatus string) error {
	err := pr.db.Model(&models.Payment{}).Where("payment_code = ?", payment_code).Update("payment_status_cd", paymentStatus).Error

	return err
}

func (pr *paymentRepository) DeletePayment(payment models.Payment) (models.Payment, error) {
	err := pr.db.Delete(&payment).Error

	return payment, err
}
