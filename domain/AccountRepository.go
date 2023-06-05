package domain

import (
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"

	"go_bank/errs"
	"go_bank/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction, a *Account) (*Transaction, *errs.AppError) {
	txn, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	result, err := txn.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		txn.Rollback()
		logger.Error("Error while insert transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	var newAmount float64
	if t.IsWithdraw() {
		newAmount = a.Amount - t.Amount
	} else {
		newAmount = a.Amount + t.Amount
	}
	_, err = txn.Exec(`UPDATE accounts SET amount = ? WHERE account_id = ?`, newAmount, t.AccountId)

	if err != nil {
		txn.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = txn.Commit()
	if err != nil {
		txn.Rollback()
		logger.Error("Error while committing transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexptect database error")
	}

	txnId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	t.TransactionId = strconv.FormatInt(txnId, 10)
	t.Amount = newAmount

	return &t, nil
}

func (d AccountRepositoryDb) FindById(id string) (*Account, *errs.AppError) {
	var a Account
	sqlSelect := "SELECT ACCOUNT_ID, CUSTOMER_ID, OPENING_DATE, ACCOUNT_TYPE, AMOUNT, STATUS FROM accounts WHERE account_id = ?"
	err := d.client.Get(&a, sqlSelect, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Account not found")
		}
		logger.Error("Error while select data from account by id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	return &a, nil

}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}
