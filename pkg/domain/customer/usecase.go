package customer

import (
	"github.com/FernandoCagale/c4-customer/pkg/entity"
)

type UseCase interface {
	Create(headers map[string]string, ecommerce *entity.Ecommerce) (err error)
	FindAll() (customers []*entity.Customer, err error)
	FindById(ID string) (customer *entity.Customer, err error)
	DeleteById(ID string) (err error)
}
