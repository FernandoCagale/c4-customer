package pkg

import (
	eventBoot "github.com/FernandoCagale/c4-customer/api/event"
	"github.com/FernandoCagale/c4-customer/api/handlers"
	"github.com/FernandoCagale/c4-customer/api/routers"
	"github.com/FernandoCagale/c4-customer/internal/event"
	"github.com/FernandoCagale/c4-customer/pkg/domain/customer"
	"github.com/google/wire"
)

var Container = wire.NewSet(customer.Set, handlers.Set, routers.Set, event.Set, eventBoot.Set)
