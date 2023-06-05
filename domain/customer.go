package domain

import (
	"go_bank/dto"
	"go_bank/errs"
)

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

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}
