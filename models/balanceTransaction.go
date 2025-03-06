package models

import "time"

type BalanceTransaction struct {
	BalanceTransactionId              int       `json:"balance_transaction_id" gorm:"primaryKey"`
	SenderUuid                        string    `json:"sender_uuid" gorm:"type:varchar(200)"`
	ReceiverUuid                      string    `json:"receiver_uuid" gorm:"type:varchar(200)"`
	BalanceTransactionCode            string    `json:"balance_transaction_code" gorm:"type:varchar(150)"`
	BalanceTransactionInAmount        string    `json:"balance_transaction_in_amount" gorm:"type:varchar(150)"`
	BalanceTransactionOutAmount       string    `json:"balance_transaction_out_amount" gorm:"type:varchar(150)"`
	BalanceTransactionDescription     string    `json:"balance_transaction_description" gorm:"type:text"`
	BalanceTransactionProcessCd       string    `json:"balance_transaction_process_cd" gorm:"type:varchar(150)"`
	BalanceTransactionStatusCd        string    `json:"balance_transaction_status_cd" gorm:"type:varchar(30)"`
	BalanceTransactionCreatedDate     time.Time `json:"balance_transaction_created_date"`
	BalanceTransactionCreatedUserUUID string    `json:"balance_transaction_created_user_uuid" gorm:"type:varchar(200)"`
	BalanceTransactionCreatedUsername string    `json:"balance_transaction_created_user_username" gorm:"type:varchar(100)"`
	BalanceTransactionUpdatedDate     time.Time `json:"balance_transaction_updated_date"`
	BalanceTransactionUpdatedUserUUID string    `json:"balance_transaction_updated_user_uuid" gorm:"type:varchar(200)"`
	BalanceTransactionUpdatedUsername string    `json:"balance_transaction_updated_user_username" gorm:"type:varchar(100)"`
}
