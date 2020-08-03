package entity

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Ecommerce struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (e Ecommerce) Validate() error {
	return validation.ValidateStruct(&e, )
}

func (e Ecommerce) ToCustomer(notify []string) Customer {
	return Customer{
		Code:   e.Code,
		Name:   e.Name,
		Email:  e.Email,
		Phone:  e.Phone,
		Notify: notify,
	}
}
