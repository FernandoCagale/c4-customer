package customer

import (
	"github.com/FernandoCagale/c4-customer/internal/errors"
	"github.com/FernandoCagale/c4-customer/internal/notify"
	"github.com/FernandoCagale/c4-customer/pkg/entity"
)

const CUSTOMER_TOPIC = "customer.registered"

type CustomerUseCase struct {
	repo     Repository
	producer Producer
	notify   notify.Notify
}

func NewUseCase(repo Repository, producer Producer, notify notify.Notify) *CustomerUseCase {
	return &CustomerUseCase{
		repo:     repo,
		producer: producer,
		notify:   notify,
	}
}

func (usecase *CustomerUseCase) FindAll() (customers []*entity.Customer, err error) {
	return usecase.repo.FindAll()
}

func (usecase *CustomerUseCase) FindById(ID string) (customer *entity.Customer, err error) {
	return usecase.repo.FindById(ID)
}

func (usecase *CustomerUseCase) DeleteById(ID string) (err error) {
	return usecase.repo.DeleteById(ID)
}

func (usecase *CustomerUseCase) Create(headers map[string]string, e *entity.Ecommerce) error {
	err := e.Validate()
	if err != nil {
		return errors.ErrInvalidPayload
	}

	notify, err := usecase.notify.GetNotify(headers)
	if err != nil {
		return err
	}

	customer := e.ToCustomer(notify)

	if err = usecase.repo.Create(&customer); err != nil {
		return err
	}

	if err := usecase.producer.Producer(CUSTOMER_TOPIC, customer); err != nil {
		return err
	}

	return nil
}
