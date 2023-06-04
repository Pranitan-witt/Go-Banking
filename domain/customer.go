package domain

import "go_bank/errs"

type Customer struct {
	Id          string `db:"CUSTOMER_ID"`
	Name        string `db:"NAME"`
	City        string `db:"CITY"`
	Zipcode     string `db:"ZIPCODE"`
	DateofBirth string `db:"DATE_OF_BIRTH"`
	Status      string `db:"STATUS"`
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
