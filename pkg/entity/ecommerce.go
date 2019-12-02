package entity

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Ecommerce struct {
	Code   string   `json:"code"`
	Name   string   `json:"name"`
	Email  string   `json:"orders"`
	Phone  string   `json:"phone"`
	Notify []string `json:"notify"`
}

func (e Ecommerce) Validate() error {
	return validation.ValidateStruct(&e,
	)
}

func (e Ecommerce) ToCustomer() Customer {
	return Customer{
		Code:   e.Code,
		Name:   e.Name,
		Email:  e.Email,
		Phone:  e.Phone,
		Notify: e.Notify,
	}
}
