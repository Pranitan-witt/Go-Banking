package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"go_bank/errs"
	"go_bank/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	byIdSql := "SELECT CUSTOMER_ID, NAME, CITY, ZIPCODE, DATE_OF_BIRTH, STATUS FROM customers WHERE CUSTOMER_ID = ?"

	var c Customer
	err := d.client.Get(&c, byIdSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while scanning customer by id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &c, nil
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var findAllSql string
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSql = "SELECT CUSTOMER_ID, NAME, CITY, ZIPCODE, DATE_OF_BIRTH, STATUS FROM customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql = "SELECT CUSTOMER_ID, NAME, CITY, ZIPCODE, DATE_OF_BIRTH, STATUS FROM customers WHERE STATUS = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// err = sqlx.StructScan(rows, &customers) // Another way when you combined with normal sql
	// if err != nil {
	// 	logger.Error("Error while scanning customers " + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected scanning error")
	// }

	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:admin@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
