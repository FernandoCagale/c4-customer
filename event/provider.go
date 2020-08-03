package event

import (
	eventImp "github.com/FernandoCagale/c4-customer/internal/event"
	"github.com/FernandoCagale/c4-customer/internal/notify"
	"github.com/FernandoCagale/c4-customer/pkg/domain/customer"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewCustomer, customer.Set, eventImp.Set, notify.Set)
