package domain

import (
	"go_bank/dto"
	"go_bank/errs"
)

type Account struct {
	AccountId   string  `db:"ACCOUNT_ID"`
	CustomerId  string  `db:"CUSTOMER_ID"`
	OpeningDate string  `db:"OPENING_DATE"`
	AccountType string  `db:"ACCOUNT_TYPE"`
	Amount      float64 `db:"AMOUNT"`
	Status      string  `db:"STATUS"`
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindById(string) (*Account, *errs.AppError)
	SaveTransaction(Transaction, *Account) (*Transaction, *errs.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount > amount
}
