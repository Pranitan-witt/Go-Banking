package domain

import (
	"go_bank/dto"
)

const WITHDRAW = "withdraw"
const DEPOSIT = "deposit"

type Transaction struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (t Transaction) IsWithdraw() bool {
	return t.TransactionType == dto.WITHDRAW
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
