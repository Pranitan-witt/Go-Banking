package dto

import (
	"strings"

	"go_bank/errs"
)

const WITHDRAW = "withdraw"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	CustomerId      string  `json:"customer_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (r TransactionRequest) IsTransactionTypeWithdraw() bool {
	// if r.TransactionType == WITHDRAW {
	// 	return true
	// }
	// return false
	return strings.ToLower(r.TransactionType) == WITHDRAW
}

func (r TransactionRequest) Validate() *errs.AppError {
	if strings.ToLower(r.TransactionType) != WITHDRAW && strings.ToLower(r.TransactionType) != DEPOSIT {
		return errs.NewValidationError("Transaction type must be withdraw or deposit")
	}

	if r.Amount < 0 {
		return errs.NewValidationError("Amount can not be less than zero")
	}

	return nil
}
