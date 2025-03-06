package services

import (
	"fmt"
	"rantaujoeang-app-backend/helpers"
	"rantaujoeang-app-backend/inputs"
	"rantaujoeang-app-backend/models"
	"rantaujoeang-app-backend/repositories"
	"strconv"
	"time"
)

type BalanceTransactionService interface {
	FindBalanceTransactionsService() ([]models.BalanceTransaction, error)
	FindBalanceTransactionByUserIdService(user_uuid string) ([]models.BalanceTransaction, error)
	FindBalanceTransactionService(balance_transaction_code string) (models.BalanceTransaction, error)
	CreateBalanceTransactionService(createBalanceTransactionInput inputs.CreateBalanceTransactionInput, currentUser map[string]string) (models.BalanceTransaction, error)
}

type balanceTransactionService struct {
	balanceTransactionRepository repositories.BalanceTransactionRepository
}

func NewBalanceTransactionService(balanceTransactionRepository repositories.BalanceTransactionRepository) *balanceTransactionService {
	return &balanceTransactionService{balanceTransactionRepository}
}

func (bts *balanceTransactionService) FindBalanceTransactionsService() ([]models.BalanceTransaction, error) {
	balanceTransactions, err := bts.balanceTransactionRepository.FindBalanceTransactions()

	return balanceTransactions, err
}

func (bts *balanceTransactionService) FindBalanceTransactionService(balance_transaction_code string) (models.BalanceTransaction, error) {
	balanceTransaction, err := bts.balanceTransactionRepository.FindBalanceTransaction(balance_transaction_code)

	return balanceTransaction, err
}

func (bts *balanceTransactionService) FindBalanceTransactionByUserIdService(user_uuid string) ([]models.BalanceTransaction, error) {
	balanceTransaction, err := bts.balanceTransactionRepository.FindBalanceTransactionByUserId(user_uuid)

	return balanceTransaction, err
}

func (bts *balanceTransactionService) CreateBalanceTransactionService(createBalanceTransactionInput inputs.CreateBalanceTransactionInput, currentUser map[string]string) (models.BalanceTransaction, error) {

	balanceTransactionCode := helpers.GenerateCode(createBalanceTransactionInput.BalanceTransactionProcessCd)

	balanceTransaction := models.BalanceTransaction{
		BalanceTransactionCode:            balanceTransactionCode,
		SenderUuid:                        createBalanceTransactionInput.BalanceTransactionSenderUuid,
		ReceiverUuid:                      createBalanceTransactionInput.BalanceTransactionReceiverUuid,
		BalanceTransactionInAmount:        createBalanceTransactionInput.BalanceTransactionInAmount,
		BalanceTransactionOutAmount:       createBalanceTransactionInput.BalanceTransactionOutAmount,
		BalanceTransactionProcessCd:       createBalanceTransactionInput.BalanceTransactionProcessCd,
		BalanceTransactionDescription:     createBalanceTransactionInput.BalanceTransactionDescription,
		BalanceTransactionStatusCd:        "sent",
		BalanceTransactionCreatedDate:     time.Now(),
		BalanceTransactionCreatedUserUUID: currentUser["user_uuid"],
		BalanceTransactionCreatedUsername: currentUser["user_username"],
	}

	newBalanceTransaction, err := bts.balanceTransactionRepository.CreateBalanceTransaction(balanceTransaction)

	newBalanceTransactionInAmount, _ := strconv.ParseFloat(createBalanceTransactionInput.BalanceTransactionInAmount, 64)
	newBalanceTransactionOutAmount, _ := strconv.ParseFloat(createBalanceTransactionInput.BalanceTransactionOutAmount, 64)

	switch createBalanceTransactionInput.BalanceTransactionProcessCd {
	case "top-up":
		checkUser, _ := bts.balanceTransactionRepository.FindUserBalanceTransaction(createBalanceTransactionInput.BalanceTransactionSenderUuid)

		newAmount, _ := strconv.ParseFloat(checkUser.UserBalanceTransactionAmount, 64)

		newAmount += newBalanceTransactionInAmount

		newAmountToString := strconv.FormatFloat(newAmount, 'f', 2, 64)

		bts.balanceTransactionRepository.UpdateUserBalanceTransaction(createBalanceTransactionInput.BalanceTransactionSenderUuid, newAmountToString)
	case "transfer":
		checkSenderUser, _ := bts.balanceTransactionRepository.FindUserBalanceTransaction(createBalanceTransactionInput.BalanceTransactionSenderUuid)
		checkReceiverUser, _ := bts.balanceTransactionRepository.FindUserBalanceTransaction(createBalanceTransactionInput.BalanceTransactionReceiverUuid)

		newSenderAmount, _ := strconv.ParseFloat(checkSenderUser.UserBalanceTransactionAmount, 64)
		newReceiverAmount, _ := strconv.ParseFloat(checkReceiverUser.UserBalanceTransactionAmount, 64)

		newSenderAmount -= newBalanceTransactionOutAmount
		newReceiverAmount += newBalanceTransactionOutAmount

		newSenderAmountToString := strconv.FormatFloat(newSenderAmount, 'f', 2, 64)
		newReceiverAmountToString := strconv.FormatFloat(newReceiverAmount, 'f', 2, 64)

		bts.balanceTransactionRepository.UpdateUserBalanceTransaction(createBalanceTransactionInput.BalanceTransactionSenderUuid, newSenderAmountToString)
		bts.balanceTransactionRepository.UpdateUserBalanceTransaction(createBalanceTransactionInput.BalanceTransactionReceiverUuid, newReceiverAmountToString)
	default:
		fmt.Println("Invalid Process !")
	}

	return newBalanceTransaction, err
}
