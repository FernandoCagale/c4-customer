//+build wireinject

package main

import (
	"github.com/FernandoCagale/c4-customer/api/routers"
	"github.com/FernandoCagale/c4-customer/internal/datastore"
	"github.com/FernandoCagale/c4-customer/pkg"
	"github.com/FernandoCagale/c4-customer/event"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func SetupApplication(*gorm.DB) (*routers.SystemRoutes, error) {
	wire.Build(pkg.Container)
	return nil, nil
}

func SetupPostgres() (*gorm.DB, error) {
	wire.Build(datastore.Set)
	return nil, nil
}

func SetupEvents(session *gorm.DB) (*event.CustomerEvent, error) {
	wire.Build(event.Set)
	return nil, nil
}