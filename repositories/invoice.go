package repositories

import (
	"rantaujoeang-app-backend/models"

	"gorm.io/gorm"
)

type InvoiceRepository interface {
	FindInvoices() ([]models.Invoice, error)
	FindInvoicesByUserId(user_uuid string) ([]models.Invoice, error)
	FindInvoice(invoice_code string) (models.Invoice, error)
	FindUser(user_uuid string) (models.User, error)
	FindDorm(dorm_code string) (models.Dorm, error)
	CreateInvoice(invoice models.Invoice) (models.Invoice, error)
	UpdateInvoice(invoice models.Invoice) (models.Invoice, error)
	UpdateDormQuantity(dorm_code string, quantityAmount string) error
	UpdateUserBalanceTransaction(user_uuid string, userBalanceAmount string) error
	UpdatePayment(payment_code string, payment_status_cd string) error
	DeleteInvoice(invoice models.Invoice) (models.Invoice, error)
}

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *invoiceRepository {
	return &invoiceRepository{db}
}

func (ir *invoiceRepository) FindInvoices() ([]models.Invoice, error) {
	var invoice []models.Invoice

	err := ir.db.Find(&invoice).Error

	return invoice, err
}

func (ir *invoiceRepository) FindInvoicesByUserId(user_uuid string) ([]models.Invoice, error) {
	var invoices []models.Invoice

	err := ir.db.Where("user_uuid = ?", user_uuid).Find(&invoices).Error

	return invoices, err
}

func (ir *invoiceRepository) FindInvoice(invoice_code string) (models.Invoice, error) {
	var invoice models.Invoice

	err := ir.db.Where("invoice_code = ?", invoice_code).First(&invoice).Error

	return invoice, err
}

func (ir *invoiceRepository) FindUser(user_uuid string) (models.User, error) {
	var user models.User

	err := ir.db.Where("user_uuid = ?", user_uuid).First(&user).Error

	return user, err
}

func (ir *invoiceRepository) FindDorm(dorm_code string) (models.Dorm, error) {
	var dorm models.Dorm

	err := ir.db.Where("dorm_code = ?").First(&dorm).Error

	return dorm, err
}

func (ir *invoiceRepository) CreateInvoice(invoice models.Invoice) (models.Invoice, error) {
	err := ir.db.Create(&invoice).Error

	return invoice, err
}

func (ir *invoiceRepository) UpdateInvoice(invoice models.Invoice) (models.Invoice, error) {
	err := ir.db.Save(&invoice).Error

	return invoice, err
}

func (ir *invoiceRepository) UpdateDormQuantity(dorm_code string, quantityAmount string) error {
	err := ir.db.Model(models.Dorm{}).Where("dorm_code = ?", dorm_code).Update("dorm_quantity_amount", quantityAmount).Error

	return err
}

func (ir *invoiceRepository) UpdateUserBalanceTransaction(user_uuid string, userBalanceAmount string) error {
	err := ir.db.Model(models.User{}).Where("user_uuid = ?", user_uuid).Update("user_balance_transaction_amount", userBalanceAmount).Error

	return err
}

func (ir *invoiceRepository) UpdatePayment(payment_code string, payment_status_cd string) error {
	err := ir.db.Model(models.Payment{}).Where("payment_code = ?", payment_code).Update("payment_status_cd", payment_status_cd).Error

	return err
}

func (ir *invoiceRepository) DeleteInvoice(invoice models.Invoice) (models.Invoice, error) {
	err := ir.db.Delete(&invoice).Error

	return invoice, err
}
