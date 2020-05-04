package event

import (
	"github.com/FernandoCagale/c4-customer/pkg/domain/customer"
	eventImp "github.com/FernandoCagale/c4-customer/internal/event"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewCustomer, customer.Set, eventImp.Set)
