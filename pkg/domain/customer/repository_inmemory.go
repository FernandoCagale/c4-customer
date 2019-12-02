package customer

import (
	"github.com/FernandoCagale/c4-customer/internal/errors"
	"github.com/FernandoCagale/c4-customer/pkg/entity"
)

type InMemoryRepository struct {
	m map[string]*entity.Customer
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{map[string]*entity.Customer{}}
}

func (repo *InMemoryRepository) FindAll() (customers []*entity.Customer, err error) {
	for _, customer := range repo.m {
		customers = append(customers, customer)
	}
	return customers, nil
}

func (repo *InMemoryRepository) FindById(ID string) (customer *entity.Customer, err error) {
	for _, customer := range repo.m {
		if customer.Code == ID {
			return customer, nil
		}
	}
	return nil, errors.ErrNotFound
}

func (repo *InMemoryRepository) DeleteById(ID string) (err error) {
	for _, customer := range repo.m {
		if customer.Code == ID {
			delete(repo.m, ID)
			return nil
		}
	}
	return errors.ErrNotFound
}

func (repo *InMemoryRepository) Create(e *entity.Customer) (err error) {
	customer := repo.m[e.Code]

	if customer == nil {
		repo.m[e.Code] = e
		return nil
	}

	return nil
}
