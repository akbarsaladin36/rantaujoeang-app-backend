package repositories

import (
	"rantaujoeang-app-backend/models"

	"gorm.io/gorm"
)

type BalanceTransactionRepository interface {
	FindBalanceTransactions() ([]models.BalanceTransaction, error)
	FindBalanceTransactionByUserId(user_uuid string) ([]models.BalanceTransaction, error)
	FindBalanceTransaction(balance_transaction_code string) (models.BalanceTransaction, error)
	CreateBalanceTransaction(balanceTransaction models.BalanceTransaction) (models.BalanceTransaction, error)
	FindUserBalanceTransaction(user_uuid string) (models.User, error)
	UpdateUserBalanceTransaction(user_uuid string, newAmount string) error
}

type balanceTransactionRepository struct {
	db *gorm.DB
}

func NewBalanceTransactionRepository(db *gorm.DB) *balanceTransactionRepository {
	return &balanceTransactionRepository{db}
}

func (btr *balanceTransactionRepository) FindBalanceTransactions() ([]models.BalanceTransaction, error) {
	var balanceTransactions []models.BalanceTransaction

	err := btr.db.Find(&balanceTransactions).Error

	return balanceTransactions, err
}

func (btr *balanceTransactionRepository) FindBalanceTransaction(balance_transaction_code string) (models.BalanceTransaction, error) {
	var balanceTransaction models.BalanceTransaction

	err := btr.db.Where("balance_transaction_code = ?", balance_transaction_code).First(&balanceTransaction).Error

	return balanceTransaction, err
}

func (btr *balanceTransactionRepository) FindBalanceTransactionByUserId(user_uuid string) ([]models.BalanceTransaction, error) {
	var balanceTransaction []models.BalanceTransaction

	err := btr.db.Where("sender_id = ?", user_uuid).Or("receiver_id = ?", user_uuid).Find(&balanceTransaction).Error

	return balanceTransaction, err
}

func (btr *balanceTransactionRepository) CreateBalanceTransaction(balanceTransaction models.BalanceTransaction) (models.BalanceTransaction, error) {
	err := btr.db.Create(&balanceTransaction).Error

	return balanceTransaction, err
}

func (btr *balanceTransactionRepository) FindUserBalanceTransaction(user_uuid string) (models.User, error) {
	var user models.User

	err := btr.db.Where("user_uuid = ?", user_uuid).First(&user).Error

	return user, err
}

func (btr *balanceTransactionRepository) UpdateUserBalanceTransaction(user_uuid string, newAmount string) error {
	err := btr.db.Model(&models.User{}).Where("user_uuid = ?", user_uuid).Update("user_balance_transaction_amount", newAmount).Error

	return err
}
