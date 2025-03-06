package responses

import (
	"rantaujoeang-app-backend/models"
	"time"
)

type BalanceTransactionRsps struct {
	BalanceTransactionCode            string    `json:"balance_transaction_code"`
	BalanceTransactionInAmount        string    `json:"balance_transaction_in_amount"`
	BalanceTransactionOutAmount       string    `json:"balance_transaction_out_amount"`
	BalanceTransactionDescription     string    `json:"balance_transaction_description"`
	BalanceTransactionProcessCd       string    `json:"balance_transaction_process_cd"`
	BalanceTransactionStatusCd        string    `json:"balance_transaction_status_cd"`
	BalanceTransactionCreatedDate     time.Time `json:"balance_transaction_created_date"`
	BalanceTransactionCreatedUserUUID string    `json:"balance_transaction_created_user_uuid"`
	BalanceTransactionCreatedUsername string    `json:"balance_transaction_created_user_username"`
}

func GetBalanceTransactionResponse(BTRsps models.BalanceTransaction) BalanceTransactionRsps {
	return BalanceTransactionRsps{
		BalanceTransactionCode:            BTRsps.BalanceTransactionCode,
		BalanceTransactionInAmount:        BTRsps.BalanceTransactionInAmount,
		BalanceTransactionOutAmount:       BTRsps.BalanceTransactionOutAmount,
		BalanceTransactionDescription:     BTRsps.BalanceTransactionDescription,
		BalanceTransactionProcessCd:       BTRsps.BalanceTransactionProcessCd,
		BalanceTransactionStatusCd:        BTRsps.BalanceTransactionStatusCd,
		BalanceTransactionCreatedDate:     BTRsps.BalanceTransactionCreatedDate,
		BalanceTransactionCreatedUserUUID: BTRsps.BalanceTransactionCreatedUserUUID,
		BalanceTransactionCreatedUsername: BTRsps.BalanceTransactionCreatedUsername,
	}
}
