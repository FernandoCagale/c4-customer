package event

import (
	"encoding/json"
	"fmt"
	"github.com/FernandoCagale/c4-customer/internal/event"
	"github.com/FernandoCagale/c4-customer/pkg/domain/customer"
	"github.com/FernandoCagale/c4-customer/pkg/entity"
	"log"
)

const (
	CUSTOMER_QUEUE = "customer-registered"
)

type CustomerEvent struct {
	usecase customer.UseCase
	event   event.Event
}

func NewCustomer(usecase customer.UseCase, event event.Event) *CustomerEvent {
	return &CustomerEvent{
		usecase: usecase,
		event:   event,
	}
}

func (eventCustomer *CustomerEvent) ProcessCustomer() {
	messages, err := eventCustomer.event.SubscribeQueue(CUSTOMER_QUEUE)
	if err != nil {
		fmt.Println(err.Error())
	}

	for msg := range messages {
		log.Printf("received message: %s, CUSTOMER: %s", msg.UUID, string(msg.Payload))

		var ecommerce entity.Ecommerce

		if err := json.Unmarshal(msg.Payload, &ecommerce); err != nil {
			fmt.Println(err.Error())
			msg.Nacked()
		}

		if err = eventCustomer.usecase.Create(&ecommerce); err != nil {
			fmt.Println(err.Error())
			msg.Nacked()
		}

		msg.Ack() //TODO x-dead-letter-exchange
	}
}
