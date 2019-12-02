package customer

import "github.com/FernandoCagale/c4-customer/pkg/entity"

type Repository interface {
	Create(customer *entity.Customer) (err error)
	FindAll() (customers []*entity.Customer, err error)
	FindById(ID string) (customer *entity.Customer, err error)
	DeleteById(ID string) (err error)
}
