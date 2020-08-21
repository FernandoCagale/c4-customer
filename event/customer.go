package event

import (
	"github.com/FernandoCagale/c4-customer/internal/consumer"
)

type CustomerEvent struct {
	consumer consumer.Consumer
}

func NewCustomer(consumer consumer.Consumer) *CustomerEvent {
	return &CustomerEvent{
		consumer: consumer,
	}
}

func (event *CustomerEvent) MakeEvents() {
	go event.consumer.Consumer()
}
