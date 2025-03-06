package inputs

type CreateBalanceTransactionInput struct {
	BalanceTransactionSenderUuid   string `json:"balance_transaction_sender_uuid"`
	BalanceTransactionReceiverUuid string `json:"balance_transaction_receiver_uuid"`
	BalanceTransactionInAmount     string `json:"balance_transaction_in_amount"`
	BalanceTransactionOutAmount    string `json:"balance_transaction_out_amount"`
	BalanceTransactionDescription  string `json:"balance_transaction_description"`
	BalanceTransactionProcessCd    string `json:"balance_transaction_process_cd"`
}
