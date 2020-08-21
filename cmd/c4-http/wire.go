//+build wireinject

package main

import (
	"github.com/FernandoCagale/c4-customer/api/routers"
	"github.com/FernandoCagale/c4-customer/event"
	"github.com/FernandoCagale/c4-customer/internal/datastore"
	"github.com/FernandoCagale/c4-customer/pkg"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func SetupApplication(*gorm.DB) (*routers.SystemRoutes, error) {
	wire.Build(pkg.ContainerHTTP)
	return nil, nil
}

func SetupPostgres() (*gorm.DB, error) {
	wire.Build(datastore.Set)
	return nil, nil
}

func SetupEvents(*gorm.DB) (*event.CustomerEvent, error) {
	wire.Build(event.SetHTTP)
	return nil, nil
}