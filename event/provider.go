package event

import (
	"github.com/FernandoCagale/c4-customer/internal/consumer"
	"github.com/FernandoCagale/c4-customer/internal/notify"
	"github.com/FernandoCagale/c4-customer/pkg/domain/customer"
	"github.com/google/wire"
)

var SetGRPC = wire.NewSet(NewCustomer, customer.Set, notify.SetGRPC, consumer.Set)

var SetHTTP = wire.NewSet(NewCustomer, customer.Set, notify.SetHTTP, consumer.Set)
