package customer

import (
	"github.com/FernandoCagale/c4-customer/internal/errors"
	"github.com/FernandoCagale/c4-customer/internal/event"
	notify "github.com/FernandoCagale/c4-customer/internal/notify"
	"github.com/FernandoCagale/c4-customer/pkg/entity"
)

const QUEUE = "notify.customer"

type CustomerUseCase struct {
	repo   Repository
	event  event.Event
	notify notify.Notify
}

func NewUseCase(repo Repository, event event.Event, notify notify.Notify) *CustomerUseCase {
	return &CustomerUseCase{
		repo:   repo,
		event:  event,
		notify: notify,
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

func (usecase *CustomerUseCase) Create(e *entity.Ecommerce) error {
	err := e.Validate()
	if err != nil {
		return errors.ErrInvalidPayload
	}

	notify, err := usecase.notify.GetNotify()
	if err != nil {
		return err
	}

	customer := e.ToCustomer(notify)

	if err = usecase.repo.Create(&customer); err != nil {
		return err
	}

	if err := usecase.event.PublishQueue(QUEUE, e); err != nil {
		return err
	}

	return nil
}
